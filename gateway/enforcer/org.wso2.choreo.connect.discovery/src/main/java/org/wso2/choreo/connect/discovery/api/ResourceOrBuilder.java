// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: wso2/discovery/api/Resource.proto

package org.wso2.choreo.connect.discovery.api;

public interface ResourceOrBuilder extends
    // @@protoc_insertion_point(interface_extends:wso2.discovery.api.Resource)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string id = 1;</code>
   * @return The id.
   */
  java.lang.String getId();
  /**
   * <code>string id = 1;</code>
   * @return The bytes for id.
   */
  com.google.protobuf.ByteString
      getIdBytes();

  /**
   * <code>string path = 2;</code>
   * @return The path.
   */
  java.lang.String getPath();
  /**
   * <code>string path = 2;</code>
   * @return The bytes for path.
   */
  com.google.protobuf.ByteString
      getPathBytes();

  /**
   * <code>repeated .wso2.discovery.api.SecurityInfo endpointSecurity = 3;</code>
   */
  java.util.List<org.wso2.choreo.connect.discovery.api.SecurityInfo> 
      getEndpointSecurityList();
  /**
   * <code>repeated .wso2.discovery.api.SecurityInfo endpointSecurity = 3;</code>
   */
  org.wso2.choreo.connect.discovery.api.SecurityInfo getEndpointSecurity(int index);
  /**
   * <code>repeated .wso2.discovery.api.SecurityInfo endpointSecurity = 3;</code>
   */
  int getEndpointSecurityCount();
  /**
   * <code>repeated .wso2.discovery.api.SecurityInfo endpointSecurity = 3;</code>
   */
  java.util.List<? extends org.wso2.choreo.connect.discovery.api.SecurityInfoOrBuilder> 
      getEndpointSecurityOrBuilderList();
  /**
   * <code>repeated .wso2.discovery.api.SecurityInfo endpointSecurity = 3;</code>
   */
  org.wso2.choreo.connect.discovery.api.SecurityInfoOrBuilder getEndpointSecurityOrBuilder(
      int index);

  /**
   * <code>map&lt;string, string&gt; security = 4;</code>
   */
  int getSecurityCount();
  /**
   * <code>map&lt;string, string&gt; security = 4;</code>
   */
  boolean containsSecurity(
      java.lang.String key);
  /**
   * Use {@link #getSecurityMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.String, java.lang.String>
  getSecurity();
  /**
   * <code>map&lt;string, string&gt; security = 4;</code>
   */
  java.util.Map<java.lang.String, java.lang.String>
  getSecurityMap();
  /**
   * <code>map&lt;string, string&gt; security = 4;</code>
   */

  java.lang.String getSecurityOrDefault(
      java.lang.String key,
      java.lang.String defaultValue);
  /**
   * <code>map&lt;string, string&gt; security = 4;</code>
   */

  java.lang.String getSecurityOrThrow(
      java.lang.String key);

  /**
   * <code>repeated string schemes = 5;</code>
   * @return A list containing the schemes.
   */
  java.util.List<java.lang.String>
      getSchemesList();
  /**
   * <code>repeated string schemes = 5;</code>
   * @return The count of schemes.
   */
  int getSchemesCount();
  /**
   * <code>repeated string schemes = 5;</code>
   * @param index The index of the element to return.
   * @return The schemes at the given index.
   */
  java.lang.String getSchemes(int index);
  /**
   * <code>repeated string schemes = 5;</code>
   * @param index The index of the value to return.
   * @return The bytes of the schemes at the given index.
   */
  com.google.protobuf.ByteString
      getSchemesBytes(int index);

  /**
   * <code>repeated string methods = 6;</code>
   * @return A list containing the methods.
   */
  java.util.List<java.lang.String>
      getMethodsList();
  /**
   * <code>repeated string methods = 6;</code>
   * @return The count of methods.
   */
  int getMethodsCount();
  /**
   * <code>repeated string methods = 6;</code>
   * @param index The index of the element to return.
   * @return The methods at the given index.
   */
  java.lang.String getMethods(int index);
  /**
   * <code>repeated string methods = 6;</code>
   * @param index The index of the value to return.
   * @return The bytes of the methods at the given index.
   */
  com.google.protobuf.ByteString
      getMethodsBytes(int index);

  /**
   * <code>repeated .wso2.discovery.api.SecurityList scopesList = 7;</code>
   */
  java.util.List<org.wso2.choreo.connect.discovery.api.SecurityList> 
      getScopesListList();
  /**
   * <code>repeated .wso2.discovery.api.SecurityList scopesList = 7;</code>
   */
  org.wso2.choreo.connect.discovery.api.SecurityList getScopesList(int index);
  /**
   * <code>repeated .wso2.discovery.api.SecurityList scopesList = 7;</code>
   */
  int getScopesListCount();
  /**
   * <code>repeated .wso2.discovery.api.SecurityList scopesList = 7;</code>
   */
  java.util.List<? extends org.wso2.choreo.connect.discovery.api.SecurityListOrBuilder> 
      getScopesListOrBuilderList();
  /**
   * <code>repeated .wso2.discovery.api.SecurityList scopesList = 7;</code>
   */
  org.wso2.choreo.connect.discovery.api.SecurityListOrBuilder getScopesListOrBuilder(
      int index);

  /**
   * <code>string tier = 8;</code>
   * @return The tier.
   */
  java.lang.String getTier();
  /**
   * <code>string tier = 8;</code>
   * @return The bytes for tier.
   */
  com.google.protobuf.ByteString
      getTierBytes();

  /**
   * <code>bool disableSecurity = 9;</code>
   * @return The disableSecurity.
   */
  boolean getDisableSecurity();

  /**
   * <code>.wso2.discovery.api.OperationPolicies policies = 10;</code>
   * @return Whether the policies field is set.
   */
  boolean hasPolicies();
  /**
   * <code>.wso2.discovery.api.OperationPolicies policies = 10;</code>
   * @return The policies.
   */
  org.wso2.choreo.connect.discovery.api.OperationPolicies getPolicies();
  /**
   * <code>.wso2.discovery.api.OperationPolicies policies = 10;</code>
   */
  org.wso2.choreo.connect.discovery.api.OperationPoliciesOrBuilder getPoliciesOrBuilder();

  /**
   * <code>.wso2.discovery.api.EndpointCluster endpoints = 11;</code>
   * @return Whether the endpoints field is set.
   */
  boolean hasEndpoints();
  /**
   * <code>.wso2.discovery.api.EndpointCluster endpoints = 11;</code>
   * @return The endpoints.
   */
  org.wso2.choreo.connect.discovery.api.EndpointCluster getEndpoints();
  /**
   * <code>.wso2.discovery.api.EndpointCluster endpoints = 11;</code>
   */
  org.wso2.choreo.connect.discovery.api.EndpointClusterOrBuilder getEndpointsOrBuilder();
}
