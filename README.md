# Rapid7 PoC Deployment Tool (InsightVM)
This tool is used by both customers and SE's alike to configure consistent [InsightVM](https://www.rapid7.com/products/insightvm/) installations for the use during poc deployments. 

## Assumptions 
1. InsightVM Security Console [installed](https://insightvm.help.rapid7.com/docs/install)
2. Root level administrative rights 

## Design Considerations
- [Single executable](https://codeburst.io/why-golang-is-great-for-portable-apps-94cf1236f481) 
- [Cross platform](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures)
- [Scalable](https://github.com/spf13/cobra) 
- Handle [static](https://github.com/bouk/staticfiles) resources easily 

##  Build instructions 
