// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by pkg/gen/guestosid. DO NOT EDIT.

package v1alpha3

// +kubebuilder:validation:Enum=dosGuest;win31Guest;win95Guest;win98Guest;winMeGuest;winNTGuest;win2000ProGuest;win2000ServGuest;win2000AdvServGuest;winXPHomeGuest;winXPProGuest;winXPPro64Guest;winNetWebGuest;winNetStandardGuest;winNetEnterpriseGuest;winNetDatacenterGuest;winNetBusinessGuest;winNetStandard64Guest;winNetEnterprise64Guest;winLonghornGuest;winLonghorn64Guest;winNetDatacenter64Guest;winVistaGuest;winVista64Guest;windows7Guest;windows7_64Guest;windows7Server64Guest;windows8Guest;windows8_64Guest;windows8Server64Guest;windows9Guest;windows9_64Guest;windows9Server64Guest;windows11_64Guest;windows12_64Guest;windowsHyperVGuest;windows2019srv_64Guest;windows2019srvNext_64Guest;windows2022srvNext_64Guest;freebsdGuest;freebsd64Guest;freebsd11Guest;freebsd11_64Guest;freebsd12Guest;freebsd12_64Guest;freebsd13Guest;freebsd13_64Guest;freebsd14Guest;freebsd14_64Guest;redhatGuest;rhel2Guest;rhel3Guest;rhel3_64Guest;rhel4Guest;rhel4_64Guest;rhel5Guest;rhel5_64Guest;rhel6Guest;rhel6_64Guest;rhel7Guest;rhel7_64Guest;rhel8_64Guest;rhel9_64Guest;centosGuest;centos64Guest;centos6Guest;centos6_64Guest;centos7Guest;centos7_64Guest;centos8_64Guest;centos9_64Guest;oracleLinuxGuest;oracleLinux64Guest;oracleLinux6Guest;oracleLinux6_64Guest;oracleLinux7Guest;oracleLinux7_64Guest;oracleLinux8_64Guest;oracleLinux9_64Guest;suseGuest;suse64Guest;slesGuest;sles64Guest;sles10Guest;sles10_64Guest;sles11Guest;sles11_64Guest;sles12Guest;sles12_64Guest;sles15_64Guest;sles16_64Guest;nld9Guest;oesGuest;sjdsGuest;mandrakeGuest;mandrivaGuest;mandriva64Guest;turboLinuxGuest;turboLinux64Guest;ubuntuGuest;ubuntu64Guest;debian4Guest;debian4_64Guest;debian5Guest;debian5_64Guest;debian6Guest;debian6_64Guest;debian7Guest;debian7_64Guest;debian8Guest;debian8_64Guest;debian9Guest;debian9_64Guest;debian10Guest;debian10_64Guest;debian11Guest;debian11_64Guest;debian12Guest;debian12_64Guest;asianux3Guest;asianux3_64Guest;asianux4Guest;asianux4_64Guest;asianux5_64Guest;asianux7_64Guest;asianux8_64Guest;asianux9_64Guest;opensuseGuest;opensuse64Guest;fedoraGuest;fedora64Guest;coreos64Guest;vmwarePhoton64Guest;other24xLinuxGuest;other26xLinuxGuest;otherLinuxGuest;other3xLinuxGuest;other4xLinuxGuest;other5xLinuxGuest;other6xLinuxGuest;genericLinuxGuest;other24xLinux64Guest;other26xLinux64Guest;other3xLinux64Guest;other4xLinux64Guest;other5xLinux64Guest;other6xLinux64Guest;otherLinux64Guest;solaris6Guest;solaris7Guest;solaris8Guest;solaris9Guest;solaris10Guest;solaris10_64Guest;solaris11_64Guest;os2Guest;eComStationGuest;eComStation2Guest;netware4Guest;netware5Guest;netware6Guest;openServer5Guest;openServer6Guest;unixWare7Guest;darwinGuest;darwin64Guest;darwin10Guest;darwin10_64Guest;darwin11Guest;darwin11_64Guest;darwin12_64Guest;darwin13_64Guest;darwin14_64Guest;darwin15_64Guest;darwin16_64Guest;darwin17_64Guest;darwin18_64Guest;darwin19_64Guest;darwin20_64Guest;darwin21_64Guest;darwin22_64Guest;darwin23_64Guest;vmkernelGuest;vmkernel5Guest;vmkernel6Guest;vmkernel65Guest;vmkernel7Guest;vmkernel8Guest;amazonlinux2_64Guest;amazonlinux3_64Guest;crxPod1Guest;rockylinux_64Guest;almalinux_64Guest;otherGuest;otherGuest64
type VirtualMachineGuestOSIdentifier string