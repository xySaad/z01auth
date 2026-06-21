# 01 Auth

### Overview

This repo includes SDK for Zone01 authentication API to simplify login process in Zone01 related platforms.

### Motivation

As a candidate at Zone01 Oujda I saw several problems related to checking integrity of users; problems such as:

- Transportation platform where poolers wheren't able to book a seat until a list of poolers is manually provided from a staff member to the transportation platform maintainer.
- Raid Audits where staff members contact talents individually to volunteer for the audit.
- Workshop booking where talents can impersonate or double book a sesssion.

Making a unifed Authentication system will solve all of those problems.

### Flow

The SDK uses the internal Gitea to retrieve the user id, then uses that id to retreive candidate informations (registred events and ban status) from the GraphQL API.

### Assumptions

- All staff members are admins in Gitea.
- All candidates registred in event module (id 41) are talents.
- All candidates registred in %piscine-go% are poolers.

### How to run

- generate graphql queries

```bash
go generate
```

- then use this package in your project
