﻿using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;

namespace Modm.Engine.Jenkins.Client
{
    public class JenkinsClientFactory
	{
        private readonly JenkinsOptions options;
        private readonly HttpClient httpClient;
        private readonly ApiTokenClient apiTokenClient;
        private readonly IServiceProvider serviceProvider;

        public JenkinsClientFactory()
        {
        }

        public JenkinsClientFactory(HttpClient client, ApiTokenClient apiTokenClient, IServiceProvider serviceProvider, IOptions<JenkinsOptions> options)
		{
            this.options = options.Value;
            this.httpClient = client;
            this.apiTokenClient = apiTokenClient;
            this.serviceProvider = serviceProvider;
        }

        public virtual async Task<IJenkinsClient> Create()
        {
            try
            {
                // to start making calls to Jenkins, an API Token is required. Fetch this token using the provider
                var apiToken = await apiTokenClient.Get();
                var jenkinsNetClient = new JenkinsNET.JenkinsClient(options.BaseUrl)
                {
                    UserName = options.UserName,
                    ApiToken = apiToken
                };

                // add the api token to the options 

                var logger = serviceProvider.GetRequiredService<ILogger<JenkinsClient>>();
                var client = new JenkinsClient(httpClient, jenkinsNetClient, logger, options);

                return client;
            }
            catch (Exception ex)
            {
                return null;
            }
            
        }
    }
}

