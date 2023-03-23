package org.example;

import org.wso2.apk.enforcer.commons.model.APIConfig;
import org.wso2.apk.enforcer.commons.model.RequestContext;
import org.wso2.apk.enforcer.commons.Filter;

import java.util.Map;

public class CustomFilter implements Filter {
    private Map<String, String> configProperties;

    @Override
    public void init(APIConfig apiConfig, Map<String, String> configProperties) {
        this.configProperties = configProperties;
    }

    @Override
    public boolean handleRequest(RequestContext requestContext) {
        String headerValue = configProperties.get("CustomProperty");
        requestContext.addOrModifyHeaders("Custom-header-1", headerValue);
        return true;
    }
}
