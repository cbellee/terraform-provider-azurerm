provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "rg" {
  name     = "${var.prefix}-adf-rg"
  location = var.location
}

resource "azurerm_data_factory" "host" {
  name                = "${var.prefix}-adf-host"
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name
}

resource "azurerm_data_factory_integration_runtime_azure" "mir" {
  name               			 = "${var.prefix}-adf-mir"
  data_factory_name  			 = azurerm_data_factory.host.name
  resource_group_name			 = azurerm_resource_group.rg.name
  compute_type 					 = "General"
  core_count					 = 8
  time_to_live_min 				 = 0
  location            			 = azurerm_resource_group.rg.location
  enable_managed_virtual_network = true
}

/* resource "azurerm_data_factory_managed_virtual_network" "managed-vnet" {
	resource_group_name = azurerm_resource_group.example.name
	data_factory_name = azurerm_data_factory.host.name
	managed_virtual_network_name = "default"
} */