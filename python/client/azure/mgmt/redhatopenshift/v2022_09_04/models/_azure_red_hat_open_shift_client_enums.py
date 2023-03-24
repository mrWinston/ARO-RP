# coding=utf-8
# --------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#   http://www.apache.org/licenses/LICENSE-2.0
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
# 
# Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code is regenerated.
# --------------------------------------------------------------------------

from enum import Enum
from six import with_metaclass
from azure.core import CaseInsensitiveEnumMeta


class CreatedByType(with_metaclass(CaseInsensitiveEnumMeta, str, Enum)):
    """The type of identity that created the resource.
    """

    USER = "User"
    APPLICATION = "Application"
    MANAGED_IDENTITY = "ManagedIdentity"
    KEY = "Key"

class EncryptionAtHost(with_metaclass(CaseInsensitiveEnumMeta, str, Enum)):
    """EncryptionAtHost represents encryption at host state
    """

    DISABLED = "Disabled"
    ENABLED = "Enabled"

class FipsValidatedModules(with_metaclass(CaseInsensitiveEnumMeta, str, Enum)):
    """FipsValidatedModules determines if FIPS is used.
    """

    DISABLED = "Disabled"
    ENABLED = "Enabled"

class ProvisioningState(with_metaclass(CaseInsensitiveEnumMeta, str, Enum)):
    """ProvisioningState represents a provisioning state.
    """

    ADMIN_UPDATING = "AdminUpdating"
    CREATING = "Creating"
    DELETING = "Deleting"
    FAILED = "Failed"
    SUCCEEDED = "Succeeded"
    UPDATING = "Updating"

class Visibility(with_metaclass(CaseInsensitiveEnumMeta, str, Enum)):
    """Visibility represents visibility.
    """

    PRIVATE = "Private"
    PUBLIC = "Public"
