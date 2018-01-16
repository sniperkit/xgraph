# 1. Insert clinic with a json file

Date: 2017-12-30

## Status

Accepted

## Context

Inserting and updating clinic using the CLI is getting cumbersome since clinic have many fields.

## Decision

* Running ./cli add-clinic clinic.json will inerts the clinic into the data store.
* Running this command will also modify an existing clinic.

## Consequences

* Insert and update clinics using the CLI will be faster.

