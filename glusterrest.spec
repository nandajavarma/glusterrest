# Determine if systemd will be used
%if ( 0%{?fedora} && 0%{?fedora} > 16 ) || ( 0%{?rhel} && 0%{?rhel} > 6 )
%global with_systemd 1
%endif

%define debug_package %{nil}
%define pkg_name glusterrest
%define pkg_version 0.0.1
%define pkg_release 1

Name: glusterfs-restapi-server
Version: %{pkg_version}
Release: %{pkg_release}%{?dist}
Summary: Modern extensible web-based storage management platform
Source0: %{pkg_name}-%{pkg_version}.tar.gz
License: MIT
Group: Applications/Storage
BuildRoot: %{_tmppath}/%{pkg_name}-%{pkg_version}-%{pkg_release}-buildroot
Url: github.com/aravindavk/glusterrest

%description
Gluster REST API server to provide Management APIs for Gluster

%if 0%{?with_systemd}
BuildRequires:  systemd
Requires(post): systemd
Requires(preun): systemd
Requires(postun): systemd
%else
Requires(post):   /sbin/chkconfig
Requires(preun):  /sbin/service
Requires(preun):  /sbin/chkconfig
Requires(postun): /sbin/service
%endif

BuildRequires: golang

%package -n glusterfs-events
Summary: Eventing for Gluster

%description -n glusterfs-events
Gluster Eventing to provide realtime notification for Gluster Events
Requires: glusterfs-restapi-server

%prep
%setup -n %{pkg_name}-%{pkg_version}

%build
make devbuild

%install
rm -rf $RPM_BUILD_ROOT
make install DESTDIR=$RPM_BUILD_ROOT PREFIX=/usr

%post

%preun

%clean
rm -rf "$RPM_BUILD_ROOT"

%files
%attr(0755, root, root) %{_sbindir}/glusterrestd
%attr(0600, root, root) %{_sysconfdir}/glusterfs/glusterrest.json
%{_var}/log/glusterfs/rest
%{_unitdir}/glusterrestd.service

%files -n glusterfs-events
%attr(0755, root, root) %{_sbindir}/glustereventsd
%{_unitdir}/glustereventsd.service


%changelog
* Thu Jan 28 2016 <avishwan@redhat.com>
- Initial Build
