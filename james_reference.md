# James Reference

```Powershell

#Variables
$sitename = read-host "What is the site name?"
$users = "internal.svp.com\grp_citrix_" + "$sitename"
$CatalogName = "MC_" + "$sitename"
$DGName = "DG_" + "$sitename"
$MachineName = "BHM-DCB-" + "$sitename"
$PIMS = read-host "What is the name of the PIMS being used? Avimark or Cornerstone (case sensitive!)"
$AppName = "$sitename" + "PIMS"
$PublishedName =  "PIMS" + "$sitename"

#Create SITE Catalog
#$brokerUsers = New-BrokerUser -Name $users
$catalog = New-BrokerCatalog -Name $CatalogName -AllocationType "Random" -Description $CatalogName -PersistUserChanges "OnLocal" -ProvisioningType "Manual" -SessionSupport "MultiSession" -MachinesArePhysical $true -ZoneUid "8bb8064d-898b-4ea9-906e-03c9bfd12c46"

#Add SITE Machine to Catalog
$BrokeredMachine = New-BrokerMachine -MachineName $MachineName -CatalogUid $catalog.uid

#Set SITE catalog VDA upgrade to LTSR only
Set-VusCatalogUpgradeType -CatalogName $CatalogName -UpgradeType LTSR


#Create new desktops & applications delivery group
$dg = New-BrokerDesktopGroup -Name $DGName -PublishedName $DGName -DesktopKind "Shared" -SessionSupport "MultiSession" -DeliveryType DesktopsAndApps -Description $DGName

#Create Appplications
if ($PIMS -eq "Avimark") {
    New-BrokerApplication -ApplicationType HostedOnDesktop -Name "$AppName" -PublishedName "$PublishedName" -MaxPerUserInstances 0 -MaxTotalInstances 0 -MaxPerMachineInstances 0 -WorkingDirectory "C:\Avimark" -CommandLineExecutable "C:\Avimark\avimark.exe" -Description "Keywords:Mandatory" -DesktopGroup $dg
  }
if ($PIMS -eq "Cornerstone") {
    New-BrokerApplication -ApplicationType HostedOnDesktop -Name "$AppName" -PublishedName "$PublishedName" -MaxPerUserInstances 0 -MaxTotalInstances 0 -MaxPerMachineInstances 0 -WorkingDirectory "C:\Cstone" -CommandLineExecutable "C:\CStone\cstone.exe" -Description "Keywords:Mandatory" -DesktopGroup $dg
  }

#Assign users to desktops and applications
New-BrokerEntitlementPolicyRule -Name $DGName -DesktopGroupUid $dg.Uid -IncludedUsers $brokerUsers -description $DGName
New-BrokerAccessPolicyRule -Name $DGName -IncludedUserFilterEnabled $true -IncludedUsers $brokerUsers -DesktopGroupUid $dg.Uid -AllowedProtocols @("HDX","RDP") 
New-BrokerAppEntitlementPolicyRule -Name $DGName -DesktopGroupUid $dg.Uid -IncludedUsers $brokerUsers -description $DGName

#Add machine to delivery group
Add-BrokerMachine -MachineName $MachineName -DesktopGroup $dg

```
