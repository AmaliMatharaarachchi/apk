package org.wso2.apk.enforcer.grpc.client;

import io.envoyproxy.envoy.extensions.common.ratelimit.v3.RateLimitDescriptor;
import io.envoyproxy.envoy.service.ratelimit.v3.RateLimitRequest;
import io.envoyproxy.envoy.service.ratelimit.v3.RateLimitServiceGrpc;
import io.envoyproxy.envoy.service.ratelimit.v3.RateLimitResponse;
import io.grpc.ManagedChannel;
import io.grpc.netty.shaded.io.grpc.netty.GrpcSslContexts;
import io.grpc.netty.shaded.io.grpc.netty.NettyChannelBuilder;
import io.grpc.netty.shaded.io.netty.handler.ssl.SslContext;
import org.wso2.apk.enforcer.config.ConfigHolder;

import java.io.File;
import java.nio.file.Paths;
import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import javax.net.ssl.SSLException;

public class RatelimitClient {
    RateLimitServiceGrpc.RateLimitServiceBlockingStub stub;
    private final ExecutorService executorService; // Add an ExecutorService field

    public RatelimitClient(){//String server, int port) {
        System.out.println("Ratelimitclient construct");
        File certFile = Paths.get(ConfigHolder.getInstance().getEnvVarConfig().getEnforcerPublicKeyPath()).toFile();
        File keyFile = Paths.get(ConfigHolder.getInstance().getEnvVarConfig().getEnforcerPrivateKeyPath()).toFile();
        SslContext sslContext = null;
        try {
            sslContext = GrpcSslContexts
                    .forClient()
                    .trustManager(ConfigHolder.getInstance().getTrustManagerFactory())
                    .keyManager(certFile, keyFile)
                    .build();
        } catch (SSLException e) {
            System.out.println("Error while generating SSL Context."+ e);
        }
        String rlHost = ConfigHolder.getInstance().getEnvVarConfig().getRatelimiterHost();
        int port = ConfigHolder.getInstance().getEnvVarConfig().getRatelimiterPort();
        System.out.println("rl host and port "+ rlHost + port);
        ManagedChannel channel = NettyChannelBuilder.forAddress(rlHost, port)
                .useTransportSecurity()
                .sslContext(sslContext)
                .build();
        this.stub = RateLimitServiceGrpc.newBlockingStub(channel);
        // Initialize the ExecutorService
        this.executorService = Executors.newFixedThreadPool(10);
    }

    public void shouldRatelimit(List<KeyValueHitsAddend> configs) {
//        System.out.println("RL task submitted");
        executorService.submit(() -> {
            System.out.println("Ratelimitclient test");
            for (KeyValueHitsAddend config : configs) {
                System.out.println("For: " + config.getKey());
                RateLimitDescriptor descriptor = RateLimitDescriptor.newBuilder()
                        .addEntries(RateLimitDescriptor.Entry.newBuilder().setKey(config.getKey()).setValue(config.getValue()).build())
                        .build();
                RateLimitRequest rateLimitRequest = RateLimitRequest.newBuilder()
                        .addDescriptors(descriptor)
                        .setDomain("Default")
                        .setHitsAddend(config.getHitsAddend())
                        .build();
                RateLimitResponse rateLimitResponse = stub.shouldRateLimit(rateLimitRequest);
                System.out.println(rateLimitResponse.getOverallCode());
            }
        });
        System.out.println("RL task submitted");

    }

    public static class KeyValueHitsAddend {
        private String key;
        private String value;
        private int hitsAddend;

        public KeyValueHitsAddend(String key, String value, int hitsAddend) {
            this.key = key;
            this.value = value;
            this.hitsAddend = hitsAddend;
        }

        public String getKey() {
            return key;
        }

        public String getValue() {
            return value;
        }

        public int getHitsAddend() {
            return hitsAddend;
        }
    }


}
