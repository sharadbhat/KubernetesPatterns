# Init Container

Enables separation of concerns by providing a separate lifecycle for initialization related tasks distinct from main container.

## Problem

Main application containers might have certain prerequisites before starting up such database schema setup, seed data installation.
