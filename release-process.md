# Release Process

This document outlines the steps involved in the release process for the NGINX Plus Go Client project.

## Versioning

The project follows [Semantic Versioning](https://semver.org/) for versioning.

## Release Planning and Development

The features that will go into the next release are reflected in the
corresponding [milestone](https://github.com/nginxinc/nginx-plus-go-client/milestones). Refer to
the [Issue Lifecycle](/ISSUE_LIFECYCLE.md) document for information on issue creation and assignment to releases.

## Releasing a New Version

1. Create an issue to define and track release-related activities. Choose a title that follows the
   format `Release X.Y.Z`.
2. Stop merging any new work into the main branch.
3. Check the release draft under the [GitHub releases](https://github.com/nginxinc/nginx-plus-go-client/releases) page
to ensure that everything is in order.
4. Create and push the release tag in the format `vX.Y.Z`:

    ```bash
    git tag -a vX.Y.Z -m "Release vX.Y.Z"
    git push origin vX.Y.Z
    ```

    As a result, the CI/CD pipeline will publish the release and announce it in the community Slack.
