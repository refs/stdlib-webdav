## Use it!

1. create an empty folder: `curl -X MKCOL 'http://localhost:8082/test'`
2. list the s&^t out of it: `url -X PROPFIND localhost:8082 -H "Depth: 1" | xmllint --format -`


```xml
<?xml version="1.0" encoding="UTF-8"?>
<D:multistatus xmlns:D="DAV:">
  <D:response>
      <D:href>/</D:href>
          <D:propstat>
	        <D:prop>
		        <D:displayname/>
			        <D:getlastmodified>Mon, 01 Feb 2021 16:20:12 GMT</D:getlastmodified>
				        <D:resourcetype>
					          <D:collection xmlns:D="DAV:"/>
						          </D:resourcetype>
							          <D:supportedlock>
								            <D:lockentry xmlns:D="DAV:">
									                <D:lockscope>
											              <D:exclusive/>
												                  </D:lockscope>
														              <D:locktype>
															                    <D:write/>
																	                </D:locktype>
																			          </D:lockentry>
																				          </D:supportedlock>
																					        </D:prop>
																						      <D:status>HTTP/1.1 200 OK</D:status>
																						          </D:propstat>
																							    </D:response>
																							      <D:response>
																							          <D:href>/test/</D:href>
																								      <D:propstat>
																								            <D:prop>
																									            <D:supportedlock>
																										              <D:lockentry xmlns:D="DAV:">
																											                  <D:lockscope>
																													                <D:exclusive/>
																															            </D:lockscope>
																																                <D:locktype>
																																		              <D:write/>
																																			                  </D:locktype>
																																					            </D:lockentry>
																																						            </D:supportedlock>
																																							            <D:resourcetype>
																																								              <D:collection xmlns:D="DAV:"/>
																																									              </D:resourcetype>
																																										              <D:displayname>test</D:displayname>
																																											              <D:getlastmodified>Mon, 01 Feb 2021 16:20:15 GMT</D:getlastmodified>
																																												            </D:prop>
																																													          <D:status>HTTP/1.1 200 OK</D:status>
																																														      </D:propstat>
																																														        </D:response>
																																															</D:multistatus>
```
