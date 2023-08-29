## Using a custom claim in the JWT token as principal

These are the steps I took to issue tokens with Azure AD with a custom claim
1. On https://portal.azure.com navigate to Azure Active Directory
2. Under "App Registrations", select "New registration"
3. Create a new app, I called it `something-that-connects-to-kafka` with the default settings.
4. Generate a cert with `openssl req -new -newkey rsa:4096 -nodes -x509 -days 365 -subj "/CN=client" -keyout key.pem -out cert.pem`
5. In the app registration navigate to "Certificates & Secrets" and "Certificates", and Upload certificate and upload the
   newly created cert.pem
6. Crate a new App registration named `kafka`
7. In the new app, registration, click "Manifest". Towards the top of the file chnage `"acceptMappedClaims": null` to `"acceptMappedC
   laims": true`
8. In the overview page, click the link below "Managed application in local directory" labelled "kafka"
9. Navigate to "Single Sign-on" then choose Edit under "Attributes & Claims"
10. Select "Add new claim" and proide your claim name, under "Source attribute" simply type in your value.
11. Select "Save"
12. Generate a jwt token by running the attached code, modifying the parameters in the main method to match
    the identities associated with the Apps and Azure AD tenant you are using.
