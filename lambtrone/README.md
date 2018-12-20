# Lambtrone

Lambtrone stands to be a daemon service that publishes the lambdas and all it's factories (everything related to the lambda stack, dependencies, environment etc) into the bucket's of all selected services to be pushed to.

## Basic Structure

The following components composes the backbone of Lambtrone :

* lambtrone-factory : Defines all low-level for delivering to the cloud vendor in question, like bucket, environment roles, workspace;
* lambtrone-flow : Most necessary steps and actions to be accomplished when some change in some other component, i.e code base updates and pushes, with final output (maybe some integration with CI/CD tools could be great);
* lambtrone-manifest : Manifests are a great way to define most of environmental and boilerplate stuff on a project. In Lambtrone, manifest will be able to estabilish the bases of the project, importing routes to specific lambda calls, security complements and bootstrapping initial setup of the stack.
