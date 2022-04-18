package six910dbi

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2020 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// go mod init github.com/Ulbora/six910-database-interface

//Six910DB Six910DB
type Six910DB interface {
	//sets security level to local or oauth using GoAuth2
	AddSecurity(s *Security) (bool, int64)
	UpdateSecurity(s *Security) bool
	GetSecurity() *Security
	DeleteSecurity() bool

	//stores
	AddStore(s *Store) (bool, int64)
	UpdateStore(s *Store) bool
	GetStore(sname string) *Store
	GetLocalStore() *Store
	GetStoreID(id int64) *Store
	GetStoreCount() int64
	GetStoreByLocal(localDomain string) *Store
	DeleteStore(id int64) bool

	//customer
	AddCustomer(c *Customer) (bool, int64)
	UpdateCustomer(c *Customer) bool
	GetCustomer(email string, storeID int64) *Customer
	GetCustomerID(id int64) *Customer
	GetCustomerUsers(cid int64, storeID int64) *[]LocalAccount
	GetCustomerList(storeID int64, start int64, end int64) *[]Customer
	DeleteCustomer(id int64) bool

	//Local Accounts when oauth not used
	AddLocalAccount(a *LocalAccount) bool
	UpdateLocalAccount(a *LocalAccount) bool
	GetLocalAccount(username string, storeID int64) *LocalAccount
	GetLocalAccountList(store int64) *[]LocalAccount
	DeleteLocalAccount(username string, storeID int64) bool

	//distributors
	AddDistributor(d *Distributor) (bool, int64)
	UpdateDistributor(d *Distributor) bool
	GetDistributor(id int64) *Distributor
	GetDistributorList(store int64) *[]Distributor
	DeleteDistributor(id int64) bool

	//Cart
	AddCart(c *Cart) (bool, int64)
	UpdateCart(c *Cart) bool
	GetCart(cid int64) *Cart
	DeleteCart(id int64) bool

	//cart item
	AddCartItem(ci *CartItem) (bool, int64)
	UpdateCartItem(ci *CartItem) bool
	GetCarItem(cartID int64, prodID int64) *CartItem
	GetCartItemList(cartID int64) *[]CartItem
	DeleteCartItem(id int64) bool

	//address
	AddAddress(a *Address) (bool, int64)
	UpdateAddress(a *Address) bool
	GetAddress(id int64) *Address
	GetAddressList(cid int64) *[]Address
	DeleteAddress(id int64) bool

	//category
	AddCategory(c *Category) (bool, int64)
	UpdateCategory(c *Category) bool
	GetCategory(id int64) *Category
	GetHierarchicalCategoryList(storeID int64) *[]Category
	GetCategoryList(storeID int64) *[]Category
	GetSubCategoryList(catID int64) *[]Category
	DeleteCategory(id int64) bool

	//shipping method
	AddShippingMethod(s *ShippingMethod) (bool, int64)
	UpdateShippingMethod(s *ShippingMethod) bool
	GetShippingMethod(id int64) *ShippingMethod
	GetShippingMethodList(storeID int64) *[]ShippingMethod
	DeleteShippingMethod(id int64) bool

	//shipping insurance
	AddInsurance(i *Insurance) (bool, int64)
	UpdateInsurance(i *Insurance) bool
	GetInsurance(id int64) *Insurance
	GetInsuranceList(storeID int64) *[]Insurance
	DeleteInsurance(id int64) bool

	//product
	AddProduct(p *Product) (bool, int64)
	UpdateProduct(p *Product) bool
	UpdateProductQuantity(p *Product) bool
	GetProductByID(id int64) *Product
	GetProductBySku(sku string, distributorID int64, storeID int64) *Product
	GetProductsByName(name string, storeID int64, start int64, end int64) *[]Product
	GetProductsByPromoted(storeID int64, start int64, end int64) *[]Product
	GetProductsByCaterory(catID int64, start int64, end int64) *[]Product
	GetProductList(storeID int64, start int64, end int64) *[]Product
	GetProductIDList(storeID int64) *[]int64
	GetProductIDListByCategories(storeID int64, catList *[]int64) *[]int64
	GetProductSubSkuList(storeID int64, parentProdID int64) *[]Product
	DeleteProduct(id int64) bool

	// //product sub sku
	// AddProductSubSku(p *ProductSubSku) (bool, int64)
	// UpdateProductSubSku(p *ProductSubSku) bool
	// UpdateProductSubSkuQuantity(p *ProductSubSku) bool
	// GetProductSubSkuByID(id int64) *ProductSubSku
	// GetProductSubSkuList(prodID int64) *[]ProductSubSku
	// DeleteProductSubSku(id int64) bool

	//product search
	GetProductManufacturerListByProductName(name string, storeID int64) *[]string
	GetProductByNameAndManufacturerName(manf string, name string, storeID int64,
		start int64, end int64) *[]Product
	GetProductManufacturerListByCatID(catID int64, storeID int64) *[]string
	GetProductByCatAndManufacturer(catID int64, manf string, storeID int64,
		start int64, end int64) *[]Product

	AddTaxRate(t *TaxRate) (bool, int64)
	UpdateTaxRate(t *TaxRate) bool
	GetTaxRate(country string, state string, storeID int64) *[]TaxRate
	GetTaxRateList(storeID int64) *[]TaxRate
	DeleteTaxRate(id int64) bool

	//Geographic Regions
	AddRegion(r *Region) (bool, int64)
	UpdateRegion(r *Region) bool
	GetRegion(id int64) *Region
	GetRegionList(storeID int64) *[]Region
	DeleteRegion(id int64) bool

	//Geographic Sub Regions
	AddSubRegion(s *SubRegion) (bool, int64)
	UpdateSubRegion(s *SubRegion) bool
	GetSubRegion(id int64) *SubRegion
	GetSubRegionList(regionID int64) *[]SubRegion
	DeleteSubRegion(id int64) bool

	//excluded sub regions
	AddExcludedSubRegion(e *ExcludedSubRegion) (bool, int64)
	UpdateExcludedSubRegion(e *ExcludedSubRegion) bool
	GetExcludedSubRegion(id int64) *ExcludedSubRegion
	GetExcludedSubRegionList(regionID int64) *[]ExcludedSubRegion
	DeleteExcludedSubRegion(id int64) bool

	//included sub regions
	AddIncludedSubRegion(e *IncludedSubRegion) (bool, int64)
	UpdateIncludedSubRegion(e *IncludedSubRegion) bool
	GetIncludedSubRegion(id int64) *IncludedSubRegion
	GetIncludedSubRegionList(regionID int64) *[]IncludedSubRegion
	DeleteIncludedSubRegion(id int64) bool

	//limit exclusions and inclusions to a zip code
	AddZoneZip(z *ZoneZip) (bool, int64)
	GetZoneZipListByExclusion(exID int64) *[]ZoneZip
	GetZoneZipListByInclusion(incID int64) *[]ZoneZip
	DeleteZoneZip(id int64) bool

	//product category
	AddProductCategory(pc *ProductCategory) bool
	GetProductCategoryList(productID int64) *[]int64
	DeleteProductCategory(pc *ProductCategory) bool

	//Orders
	AddOrder(o *Order) (bool, int64)
	UpdateOrder(o *Order) bool
	GetOrder(id int64) *Order
	GetOrderList(cid int64, storeID int64) *[]Order
	GetStoreOrderList(storeID int64) *[]Order
	GetStoreOrderListByStatus(status string, storeID int64) *[]Order
	GetOrderCountData(storeID int64) *[]OrderCountData
	GetOrderSalesData(storeID int64) *[]OrderSalesData
	DeleteOrder(id int64) bool

	//Visitors
	AddVisit(v *Visitor) bool
	GetVisitorData(storeID int64) *[]VisitorData

	//Order Items
	AddOrderItem(i *OrderItem) (bool, int64)
	UpdateOrderItem(i *OrderItem) bool
	GetOrderItem(id int64) *OrderItem
	GetOrderItemList(orderID int64) *[]OrderItem
	DeleteOrderItem(id int64) bool

	//Order Comments
	AddOrderComments(c *OrderComment) (bool, int64)
	GetOrderCommentList(orderID int64) *[]OrderComment

	//Order Payment Transactions
	AddOrderTransaction(t *OrderTransaction) (bool, int64)
	GetOrderTransactionList(orderID int64) *[]OrderTransaction

	//shipment
	AddShipment(s *Shipment) (bool, int64)
	UpdateShipment(s *Shipment) bool
	GetShipment(id int64) *Shipment
	GetShipmentList(orderID int64) *[]Shipment
	DeleteShipment(id int64) bool

	//shipment boxes
	AddShipmentBox(sb *ShipmentBox) (bool, int64)
	UpdateShipmentBox(sb *ShipmentBox) bool
	GetShipmentBox(id int64) *ShipmentBox
	GetShipmentBoxList(shipmentID int64) *[]ShipmentBox
	DeleteShipmentBox(id int64) bool

	//Shipment Items in box
	AddShipmentItem(si *ShipmentItem) (bool, int64)
	UpdateShipmentItem(si *ShipmentItem) bool
	GetShipmentItem(id int64) *ShipmentItem
	GetShipmentItemList(shipmentID int64) *[]ShipmentItem
	GetShipmentItemListByBox(boxNumber int64, shipmentID int64) *[]ShipmentItem
	DeleteShipmentItem(id int64) bool

	//Global Plugins
	AddPlugin(p *Plugins) (bool, int64)
	UpdatePlugin(p *Plugins) bool
	GetPlugin(id int64) *Plugins
	GetPluginList(start int64, end int64) *[]Plugins
	DeletePlugin(id int64) bool

	//store plugins installed
	AddStorePlugin(sp *StorePlugins) (bool, int64)
	UpdateStorePlugin(sp *StorePlugins) bool
	GetStorePlugin(id int64) *StorePlugins
	GetStorePluginList(storeID int64) *[]StorePlugins
	DeleteStorePlugin(id int64) bool

	//Plugins that are payment gateways
	AddPaymentGateway(pgw *PaymentGateway) (bool, int64)
	UpdatePaymentGateway(pgw *PaymentGateway) bool
	GetPaymentGateway(id int64) *PaymentGateway
	GetPaymentGatewayByName(name string, storeID int64) *PaymentGateway
	GetPaymentGateways(storeID int64) *[]PaymentGateway
	DeletePaymentGateway(id int64) bool

	//store shipment carrier like UPS and FEDex
	AddShippingCarrier(c *ShippingCarrier) (bool, int64)
	UpdateShippingCarrier(c *ShippingCarrier) bool
	GetShippingCarrier(id int64) *ShippingCarrier
	GetShippingCarrierList(storeID int64) *[]ShippingCarrier
	DeleteShippingCarrier(id int64) bool

	//----UI Cluster installation: this is only called if UI is running in a cluster---
	//Handle the situation where clients are running in a cluster
	//This gives a way to make sure the json datastores are update on each node in the cluster

	//----------------start datastore------------------------------------
	//this gets called when a node start up and add only if it doesn't already exist
	AddLocalDatastore(d *LocalDataStore) bool

	//This get get called when a change is made to a datastore from a node in the cluster
	//Or after all reloads have completed and then get set to Reload = false
	UpdateLocalDatastore(d *LocalDataStore) bool

	//This gets call by cluster nodes to see if there are pending reload
	GetLocalDatastore(storeID int64, dataStoreName string) *LocalDataStore

	//---------------------start instance--------------------
	// this gets called when each instance is started and added only if never added before
	//The instance name is pulled from Docker or an manually entered env variable
	AddInstance(i *Instances) bool

	//This gets called after instance gets reloaded
	UpdateInstance(i *Instances) bool

	//Gets called before updating an instance
	GetInstance(name string, dataStoreName string, storeID int64) *Instances

	//Gets called before updating an instance or after to see if there are any instances left to update
	GetInstanceList(dataStoreName string, storeID int64) *[]Instances

	//-------------------start write lock-------------
	//gets called after UI makes changes to one of the datastores
	//If the datastore already exists, the Update method is called from within add
	AddDataStoreWriteLock(w *DataStoreWriteLock) bool
	UpdateDataStoreWriteLock(w *DataStoreWriteLock) bool

	//gets called from within the add method and by any node trying to update a datastore
	GetDataStoreWriteLock(dataStore string, storeID int64) *DataStoreWriteLock
}
