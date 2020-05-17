# melody
Systems Monitoring https://melody.systems

Repository structure:
- `agent/`
  - The system metric collection agent application. Contains code which interacts with all sorts of system OS tools for data gathering. The customer will have access to the agent binary. The agent should be activated with a customer code so it is authorized to talk to the API. Agent sends collected metrics to the API server.
- `api/` - `https://api.melody.systems`
  - The webserver which handles central receiving, processing, and sending of any data. Should also handle user login authorization. Agents send raw metric data here, and the dashboard pulls data back out from here. Runs on our end.
- `dashboard/` - `https://dashboard.melody.systems`
  - The webserver which handles visualization (charts and trends) of data for the user. The server runs on our end but its content is accessible by the user. In theory, the dashboard server should be the only one making data `GET` calls to the API.
- `web/` - `https://melody.systems`
  - The webserver for hosting general infomation static content and front-end login pages. 
