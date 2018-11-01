# CSI lib-fc 

A go package that can be imported to help CSI plugins with connecting to fibre channel devices. Simply import it and get access to the necessary functions

## Goals

Provide a basic, lightweight library for CSI Plugin Authors to leverage some of the common tasks on a node like connecting and disconnecting fibre channel devices. 

We intentionally avoid pulling in additional dependencies, and we intend to be stateless and as such are not using receivers.  Currently the focus is strictly based on a CSI context.

## Design Philosophy
 The idea is to keep this as lightweight and generic as possible.  We intentionally avoid the use of any third party
libraries or packages in this project.  We don't have a vendor directory, because we attempt to rely only on the std
golang libs.  This may prove to not be ideal, and may be changed over time, but initially it's a worthwhile goal. 

## Community, discussion, contribution, and support

Learn how to engage with the Kubernetes community on the [community page](http://kubernetes.io/community/).

You can reach the maintainers of this project at:

- [Slack](http://slack.k8s.io/)
  * sig-storage
  * wg-csi
- [Mailing List](https://groups.google.com/forum/#!forum/kubernetes-dev)

### Code of conduct

Participation in the Kubernetes community is governed by the [Kubernetes Code of Conduct](code-of-conduct.md).

[owners]: https://git.k8s.io/community/contributors/guide/owners.md
[Creative Commons 4.0]: https://git.k8s.io/website/LICENSE
