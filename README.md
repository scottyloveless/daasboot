# daasboot

## Usage

1. Clone this repo
2. Create a .env file in the root of the project with the following format and add the correct information:

```
CLIENT_ID="API client ID"
CLIENT_SECRET="API secret"
CCID="Citrix CCID"
SITE_ID="Citrix Site_ID"
ORG_ID="Citrix Organization ID"
```

3. Build the project with `go build -o daasboot`
4. Run it with the following flags

```
./daasboot --clinic CLINICCODE --pims Avimark/Cstone
```

