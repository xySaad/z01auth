# 01 Auth

### Overview

This repo includes SDK for 01-edu authentication API to simplify login process in 01-edu related platforms.

### Motivation

As a talent (student) at Zone01Oujda I saw several problems related to checking integrity of users; problems such as:

- Transportation platform where poolers wheren't able to book a seat until a list of poolers is manually provided from staff members to the transportation platform maintainer.
- Raid Audit where staff members write an announcement to notify talents that there's an upcoming raid audit to volunteer as auditors.
- Workshop booking where talents can book in the name of other talents or double book a sesssion.

Making a unifed Authentication system will solve all of those problems.

### Flow

The SDK uses the internal Gitea to retrieve the candidate login, then uses the login to retreive registred events about the candidate from the GraphQL API.

The SDK supports authentication for poolers, talents and staff members.

For staff member it checks if the user is a Gitea admin.

### Assumptions

- Gitea's login is unique, can't be changed and is the same as the GraphQL login.
- All banned members has banned state in GraphQL.
- All staff members are admins in Gitea.
- All candidates registred in Module event are talents.
- All candidates registred in piscine-go are poolers.
