# Deploy to Defang: Remix Notetaking App

### Requirements

To run locally, you'll need 
- [Docker](https://docs.docker.com/get-docker/) to run a Postgres database
- [Node.js](https://nodejs.org/en/download/) to run the app

To deploy to Defang, you'll need
- [Pulumi CLI](https://www.pulumi.com/docs/get-started/install/) to deploy the app
- [Defang CLI](https://github.com/defang-io/defang/releases/) to deploy to Defang
- [Aiven](https://console.aiven.io/signup) account to create a Postgres database

> **Note:** You'll need to make sure, in your Aiven account, that you've got an organization and a billing group with a valid payment method.


## Get up and running locally

The following will create a temporary Postgres database (it will not persist data once you stop it)

```
docker run --name defang-remix-db --rm -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
```

Next, cd into the remix directory, copy the sample env file and install dependencies

```
cd remix
cp .env.sample .env
npm install
```

Then, run the migrations and start the app by running

```
npm run dev
```

You can now access the app at http://localhost:3000

### Stopping the database

To stop the database, run

```
docker stop defang-remix-db
```

## Deploying

cd into the pulumi directory

```
cd pulumi # or cd ../pulumi if you're still in the remix directory
```

Then, make sure you're logged into Defang and Pulumi.

```
defang login
pulumi login
```

Install the dependencies.

```
npm install
```

Initialize a stack:

```
pulumi stack init dev
```

Set the following config values:

```
pulumi config set --secret aiven:apiToken YOUR_AIVEN_AUTH_TOKEN
pulumi config set --secret aivenOrganizationId YOUR_AIVEN_ORG_ID
pulumi config set --secret aivenBillingGroupId YOUR_AIVEN_BILLING_GROUP_ID
```

Run `pulumi up` to preview and deploy changes.

You're all set! You can now run `defang ls` to see your app details and access it at the URL provided.

