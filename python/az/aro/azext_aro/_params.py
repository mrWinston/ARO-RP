# Copyright (c) Microsoft Corporation.
# Licensed under the Apache License 2.0.

from azext_aro._actions import AROPlatformWorkloadIdentityAddAction
from azext_aro._validators import validate_cidr
from azext_aro._validators import validate_client_id
from azext_aro._validators import validate_client_secret
from azext_aro._validators import validate_cluster_resource_group
from azext_aro._validators import validate_disk_encryption_set
from azext_aro._validators import validate_domain
from azext_aro._validators import validate_pull_secret
from azext_aro._validators import validate_subnet
from azext_aro._validators import validate_visibility
from azext_aro._validators import validate_vnet
from azext_aro._validators import validate_vnet_resource_group_name
from azext_aro._validators import validate_worker_count
from azext_aro._validators import validate_worker_vm_disk_size_gb
from azext_aro._validators import validate_refresh_cluster_credentials
from azext_aro._validators import validate_version_format
from azext_aro._validators import validate_outbound_type
from azext_aro._validators import validate_load_balancer_managed_outbound_ip_count
from azext_aro._validators import validate_enable_managed_identity
from azext_aro._validators import validate_platform_workload_identities
from azure.cli.core.commands.parameters import name_type
from azure.cli.core.commands.parameters import get_enum_type, get_three_state_flag
from azure.cli.core.commands.parameters import resource_group_name_type
from azure.cli.core.commands.parameters import tags_type
from azure.cli.core.commands.validators import get_default_location_from_resource_group


def load_arguments(self, _):
    with self.argument_context('aro') as c:
        c.argument('location',
                   validator=get_default_location_from_resource_group)
        c.argument('resource_name',
                   name_type,
                   help='Name of cluster.')
        c.argument('tags',
                   tags_type)

        c.argument('pull_secret',
                   help='Pull secret of cluster.',
                   validator=validate_pull_secret)
        c.argument('domain',
                   help='Domain of cluster.',
                   validator=validate_domain)
        c.argument('cluster_resource_group',
                   help='Resource group of cluster.',
                   validator=validate_cluster_resource_group)
        c.argument('fips_validated_modules', arg_type=get_three_state_flag(),
                   options_list=['--fips-validated-modules', '--fips'],
                   help='Use FIPS validated cryptography modules. [Default: false]')

        c.argument('client_id',
                   help='Client ID of cluster service principal.',
                   validator=validate_client_id)
        c.argument('client_secret',
                   help='Client secret of cluster service principal.',
                   validator=validate_client_secret(isCreate=True))

        c.argument('version',
                   options_list=['--version', c.deprecate(target='--install-version', redirect='--version', hide=True)],
                   help='OpenShift version to use for cluster creation.',
                   validator=validate_version_format)

        c.argument('pod_cidr',
                   help='CIDR of pod network. Must be a minimum of /18 or larger. [Default: 10.128.0.0/14]',
                   validator=validate_cidr('pod_cidr'))
        c.argument('service_cidr',
                   help='CIDR of service network. Must be a minimum of /18 or larger. [Default: 172.30.0.0/16]',
                   validator=validate_cidr('service_cidr'))
        c.argument('outbound_type',
                   help='Outbound type of cluster. Must be "Loadbalancer" or "UserDefinedRouting". \
                   [Default: Loadbalancer]',
                   validator=validate_outbound_type)
        c.argument('enable_preconfigured_nsg', arg_type=get_three_state_flag(),
                   help='Use Preconfigured NSGs. Allowed values: false, true. [Default: false]')
        c.argument('disk_encryption_set',
                   help='ResourceID of the DiskEncryptionSet to be used for master and worker VMs.',
                   validator=validate_disk_encryption_set)
        c.argument('master_encryption_at_host', arg_type=get_three_state_flag(),
                   options_list=['--master-encryption-at-host', '--master-enc-host'],
                   help='Encryption at host flag for master VMs. [Default: false]')
        c.argument('master_vm_size',
                   help='Size of master VMs. [Default: Standard_D8s_v3]')

        c.argument('worker_encryption_at_host', arg_type=get_three_state_flag(),
                   options_list=['--worker-encryption-at-host', '--worker-enc-host'],
                   help='Encryption at host flag for worker VMs. [Default: false]')
        c.argument('worker_vm_size',
                   help='Size of worker VMs. [Default: Standard_D4s_v3]')
        c.argument('worker_vm_disk_size_gb',
                   type=int,
                   help='Disk size in GB of worker VMs. [Default: 128]',
                   validator=validate_worker_vm_disk_size_gb)
        c.argument('worker_count',
                   type=int,
                   help='Count of worker VMs. [Default: 3]',
                   validator=validate_worker_count)

        c.argument('apiserver_visibility', arg_type=get_enum_type(['Private', 'Public']),
                   help='API server visibility. [Default: Public]',
                   validator=validate_visibility('apiserver_visibility'))

        c.argument('ingress_visibility', arg_type=get_enum_type(['Private', 'Public']),
                   help='Ingress visibility. [Default: Public]',
                   validator=validate_visibility('ingress_visibility'))

        c.argument('vnet_resource_group_name',
                   resource_group_name_type,
                   options_list=['--vnet-resource-group'],
                   help='Name of vnet resource group.',
                   validator=validate_vnet_resource_group_name)
        c.argument('vnet',
                   help='Name or ID of vnet.  If name is supplied, `--vnet-resource-group` must be supplied.',
                   validator=validate_vnet)
        c.argument('master_subnet',
                   help='Name or ID of master vnet subnet.  If name is supplied, `--vnet` must be supplied.',
                   validator=validate_subnet('master_subnet'))
        c.argument('worker_subnet',
                   help='Name or ID of worker vnet subnet.  If name is supplied, `--vnet` must be supplied.',
                   validator=validate_subnet('worker_subnet'))
        c.argument('load_balancer_managed_outbound_ip_count',
                   type=int,
                   help='The desired number of IPv4 outbound IPs created and managed by Azure for the cluster public load balancer.',  # pylint: disable=line-too-long
                   validator=validate_load_balancer_managed_outbound_ip_count,
                   options_list=['--load-balancer-managed-outbound-ip-count', '--lb-ip-count'])

        c.argument('enable_managed_identity', arg_group='Identity', arg_type=get_three_state_flag(),
                   options_list=['--enable-managed-identity', '--enable-mi'],
                   validator=validate_enable_managed_identity,
                   help='Enable managed identity for this cluster.', is_preview=True)
        c.argument('platform_workload_identities', arg_group='Identity',
                   help='Assign a platform workload identity used within the cluster', is_preview=True,
                   options_list=['--assign-platform-workload-identity', '--assign-platform-wi'],
                   validator=validate_platform_workload_identities,
                   action=AROPlatformWorkloadIdentityAddAction, nargs='+')

    with self.argument_context('aro update') as c:
        c.argument('client_secret',
                   help='Client secret of cluster service principal.',
                   validator=validate_client_secret(isCreate=False))
        c.argument('refresh_cluster_credentials',
                   arg_type=get_three_state_flag(),
                   help='Refresh cluster application credentials.',
                   options_list=['--refresh-credentials'],
                   validator=validate_refresh_cluster_credentials)

    with self.argument_context('aro get-admin-kubeconfig') as c:
        c.argument('file',
                   help='Path to the file where kubeconfig should be saved. Default: kubeconfig in local directory',
                   options_list=['--file', '-f'])
