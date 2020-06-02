package six910dbi

import "time"

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

//Security Security
type Security struct {
	ID      int64 `json:"id"`
	OauthOn bool  `json:"oauthOn"`
}

//Store Store
type Store struct {
	ID            int64     `json:"id"`
	Company       string    `json:"company"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	LocalDomain   string    `json:"localDomain"`
	RemoteDomain  string    `json:"remoteDomain"`
	OauthClientID int64     `json:"oauthClientId"`
	OauthSecret   string    `json:"oauthSecret"`
	Email         string    `json:"email"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	DateEntered   time.Time `json:"entered"`
	DateUpdated   time.Time `json:"updated"`
	StoreName     string    `json:"storeName"`
	StoreSlogan   string    `json:"storeSlogan"`
	Logo          string    `json:"logo"`
	Currency      string    `json:"currency"`
	Enabled       bool      `json:"enabled"`
}

//Customer Customer
type Customer struct {
	ID            int64     `json:"id"`
	Email         string    `json:"email"`
	ResetPassword bool      `json:"resetPassword"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Company       string    `json:"company"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Zip           string    `json:"zip"`
	Phone         string    `json:"phone"`
	DateEntered   time.Time `json:"entered"`
	DateUpdated   time.Time `json:"updated"`
	StoreID       int64     `json:"storeId"`
}

//LocalAccount LocalAccount
type LocalAccount struct {
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	Enabled     bool      `json:"enabled"`
	Role        string    `json:"role"`
	DateEntered time.Time `json:"entered"`
	DateUpdated time.Time `json:"updated"`
	StoreID     int64     `json:"storeId"`
	CustomerID  int64     `json:"customerId"`
}

//Distributor Distributor
type Distributor struct {
	ID          int64  `json:"id"`
	Company     string `json:"company"`
	ContactName string `json:"contactName"`
	Phone       string `json:"phone"`
	StoreID     int64  `json:"storeId"`
}

//Cart Cart
type Cart struct {
	ID          int64     `json:"id"`
	StoreID     int64     `json:"storeId"`
	CustomerID  int64     `json:"customerId"`
	DateEntered time.Time `json:"entered"`
	DateUpdated time.Time `json:"updated"`
}

//CartItem CartItem
type CartItem struct {
	ID        int64 `json:"id"`
	Quantity  int64 `json:"quantity"`
	CartID    int64 `json:"cartId"`
	ProductID int64 `json:"productId"`
}

//Address Address
type Address struct {
	ID         int64  `json:"id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zip        string `json:"zip"`
	County     string `json:"county"`
	Country    string `json:"country"`
	Type       string `json:"type"`
	CustomerID int64  `json:"customerId"`
}

//Category Category
type Category struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"desc"`
	Image            string `json:"image"`
	Thumbnail        string `json:"thumbnail"`
	StoreID          int64  `json:"storeId"`
	ParentCategoryID int64  `json:"parentCategoryId"`
}

//ShippingCarrier ShippingCarrier
type ShippingCarrier struct {
	ID      int64  `json:"id"`
	Carrier string `json:"carrier"`
	Type    string `json:"type"`
	StoreID int64  `json:"storeId"`
}

//ShippingMethod ShippingMethod
type ShippingMethod struct {
	ID                int64   `json:"id"`
	Name              string  `json:"name"`
	Cost              float64 `json:"cost"`
	MaxWeight         int64   `json:"maxWeight"`
	Handling          float64 `json:"handling"`
	MinOrderAmount    float64 `json:"minOrderAmount"`
	MaxOrderAmount    float64 `json:"maxOrderAmount"`
	StoreID           int64   `json:"storeId"`
	RegionID          int64   `json:"regionId"`
	ShippingCarrierID int64   `json:"shippingCarrierId"`
	InsuranceID       int64   `json:"insuranceId"`
}

//Insurance Insurance
type Insurance struct {
	ID             int64   `json:"id"`
	Cost           float64 `json:"cost"`
	MinOrderAmount float64 `json:"minOrderAmount"`
	MaxOrderAmount float64 `json:"maxOrderAmount"`
	StoreID        int64   `json:"storeId"`
}

//Product Product
type Product struct {
	ID                    int64     `json:"id"`
	Sku                   string    `json:"sku"`
	Gtin                  string    `json:"gtin"`
	Name                  string    `json:"name"`
	ShortDesc             string    `json:"shortDesc"`
	Desc                  string    `json:"desc"`
	Cost                  float64   `json:"cost"`
	Msrp                  float64   `json:"msrp"`
	Map                   float64   `json:"map"`
	Price                 float64   `json:"price"`
	SalePrice             float64   `json:"salePrice"`
	Currency              string    `json:"currency"`
	Manufacturer          string    `json:"manufacturer"`
	Stock                 int64     `json:"stock"`
	StockAlert            int64     `json:"stockAlert"`
	Weight                float64   `json:"weight"`
	Width                 float64   `json:"width"`
	Height                float64   `json:"height"`
	Depth                 float64   `json:"depth"`
	ShippingMarkup        float64   `json:"shippingMarkup"`
	Visible               bool      `json:"visible"`
	Searchable            bool      `json:"searchable"`
	MultiBox              bool      `json:"multibox"`
	ShipSeperately        bool      `json:"shipSeperately"`
	FreeShipping          bool      `json:"freeShipping"`
	Promoted              bool      `json:"promoted"`
	Dropship              bool      `json:"dropship"`
	SpecialProcessing     bool      `json:"specialProcessing"`
	SpecialProcessingType string    `json:"specialProcessingType"`
	Size                  string    `json:"size"`
	Color                 string    `json:"color"`
	Thumbnail             string    `json:"thumbnail"`
	Image1                string    `json:"image1"`
	Image2                string    `json:"image2"`
	Image3                string    `json:"image3"`
	Image4                string    `json:"image4"`
	DistributorID         int64     `json:"distributorId"`
	StoreID               int64     `json:"storeId"`
	ParentProductID       int64     `json:"parentProductId"`
	DateEntered           time.Time `json:"entered"`
	DateUpdated           time.Time `json:"updated"`
}

//Region Region
type Region struct {
	ID         int64  `json:"id"`
	RegionCode string `json:"regionCode"`
	Name       string `json:"name"`
	StoreID    int64  `json:"storeId"`
}

//SubRegion SubRegion
type SubRegion struct {
	ID            int64  `json:"id"`
	SubRegionCode string `json:"subRegionCode"`
	Name          string `json:"name"`
	RegionID      int64  `json:"regionId"`
}

//ExcludedSubRegion ExcludedSubRegion
type ExcludedSubRegion struct {
	ID               int64 `json:"id"`
	RegionID         int64 `json:"regionId"`
	ShippingMethodID int64 `json:"shippingMethodId"`
	SubRegionID      int64 `json:"subRegionId"`
}

//IncludedSubRegion IncludedSubRegion
type IncludedSubRegion struct {
	ID               int64 `json:"id"`
	RegionID         int64 `json:"regionId"`
	ShippingMethodID int64 `json:"shippingMethodId"`
	SubRegionID      int64 `json:"subRegionId"`
}

//ZoneZip ZoneZip
type ZoneZip struct {
	ID                  int64  `json:"id"`
	ZipCode             string `json:"zipCode"`
	IncludedSubRegionID int64  `json:"includedSubRegionId"`
	ExcludedSubRegionID int64  `json:"excludedSubRegionId"`
}

//ProductCategory ProductCategory
type ProductCategory struct {
	CategoryID int64 `json:"categoryId"`
	ProductID  int64 `json:"productId"`
}

//Plugins Plugins
type Plugins struct {
	ID               int64   `json:"id"`
	PluginName       string  `json:"pluginName"`
	Developer        string  `json:"developer"`
	ContactPhone     string  `json:"contactPhone"`
	DeveloperAddress string  `json:"developerAddress"`
	Fee              float64 `json:"fee"`
	Enabled          bool    `json:"enabled"`
	Category         string  `json:"category"`
	ActivateURL      string  `json:"activateUrl"`
	OauthRedirectURL string  `json:"oauthRedirectUrl"`
	IsPGW            bool    `json:"isPgw"`
}

//StorePlugins StorePlugins
type StorePlugins struct {
	ID               int64     `json:"id"`
	PluginName       string    `json:"pluginName"`
	Category         string    `json:"category"`
	Active           bool      `json:"active"`
	OauthClientID    int64     `json:"oauthClientId"`
	OauthSecret      string    `json:"oauthSecret"`
	OauthRedirectURL string    `json:"oauthRedirectUrl"`
	ActivateURL      string    `json:"activateUrl"`
	APIKey           string    `json:"apiKey"`
	RekeyTryCount    int64     `json:"rekeyTryCount"`
	RekeyDate        time.Time `json:"rekeyDate"`
	IframeURL        string    `json:"iframeUrl"`
	MenuTitle        string    `json:"menuTitle"`
	MenuIconURL      string    `json:"menuIconUrl"`
	IsPGW            bool      `json:"isPgw"`
	PluginID         int64     `json:"pluginId"`
	StoreID          int64     `json:"storeId"`
}

//PaymentGateway PaymentGateway
type PaymentGateway struct {
	ID             int64  `json:"id"`
	CheckoutURL    string `json:"checkoutUrl"`
	PostOrderURL   string `json:"postOrderUrl"`
	LogoURL        string `json:"logoUrl"`
	ClientID       string `json:"clientId"`
	ClientKey      string `json:"clientKey"`
	StorePluginsID int64  `json:"storePluginsId"`
}

//Shipment Shipment
type Shipment struct {
	ID               int64     `json:"id"`
	Status           string    `json:"status"`
	Boxes            int64     `json:"boxes"`
	ShippingHandling float64   `json:"shippingHandling"`
	Insurance        float64   `json:"insurance"`
	CreateDate       time.Time `json:"createDate"`
	Updated          time.Time `json:"updated"`
	OrderID          int64     `json:"orderId"`
}

//ShipmentBox ShipmentBox
type ShipmentBox struct {
	ID                int64     `json:"id"`
	BoxNumber         int64     `json:"boxNumber"`
	Dropship          bool      `json:"dropship"`
	Cost              float64   `json:"cost"`
	Insurance         float64   `json:"insurance"`
	Weight            float64   `json:"weight"`
	Width             float64   `json:"width"`
	Height            float64   `json:"height"`
	Depth             float64   `json:"depth"`
	TrackingNumber    string    `json:"trackingNumber"`
	ShippingAddress   string    `json:"shippingAddress"`
	ShipDate          time.Time `json:"shipDate"`
	Updated           time.Time `json:"updated"`
	ShippingMethodID  int64     `json:"shippingMethodId"`
	ShippingAddressID int64     `json:"shippingAddressId"`
	ShipmentID        int64     `json:"shipmentId"`
}

//ShipmentItem ShipmentItem
type ShipmentItem struct {
	ID            int64     `json:"id"`
	Quantity      int64     `json:"quantity"`
	Updated       time.Time `json:"updated"`
	OrderItemID   int64     `json:"orderItemId"`
	ShipmentBoxID int64     `json:"shipmentBoxId"`
	ShipmentID    int64     `json:"shippingId"`
}

//Order Order
type Order struct {
	ID                int64     `json:"id"`
	OrderDate         time.Time `json:"orderDate"`
	Updated           time.Time `json:"updated"`
	Status            string    `json:"status"`
	Subtotal          float64   `json:"subTotal"`
	ShippingHandling  float64   `json:"shippingHandling"`
	Insurance         float64   `json:"insurance"`
	Taxes             float64   `json:"taxes"`
	Total             float64   `json:"total"`
	OrderNumber       string    `json:"orderNumber"`
	OrderType         string    `json:"orderType"`
	Pickup            bool      `json:"pickup"`
	Username          string    `json:"username"`
	CustomerName      string    `json:"customerName"`
	CustomerID        int64     `json:"customerId"`
	BillingAddress    string    `json:"billingAddress"`
	BillingAddressID  int64     `json:"billingAddressId"`
	ShippingAddress   string    `json:"shippingAddress"`
	ShippingAddressID int64     `json:"shippingAddressId"`
	StoreID           int64     `json:"storeId"`
}

//OrderItem OrderItem
type OrderItem struct {
	ID               int64  `json:"id"`
	Quantity         int64  `json:"quantity"`
	BackOrdered      bool   `json:"backOrdered"`
	Dropship         bool   `json:"dropship"`
	ProductName      string `json:"productName"`
	ProductShortDesc string `json:"productShortDesc"`
	ProductID        int64  `json:"productId"`
	OrderID          int64  `json:"orderId"`
}

//OrderComment OrderComment
type OrderComment struct {
	ID       int64  `json:"id"`
	Comment  string `json:"comment"`
	Username string `json:"username"`
	OrderID  int64  `json:"orderId"`
}

//OrderTransaction OrderTransaction
type OrderTransaction struct {
	ID              int64     `json:"id"`
	DateEntered     time.Time `json:"entered"`
	TransactionID   string    `json:"transactionId"`
	Type            string    `json:"type"`
	Method          string    `json:"method"`
	Amount          float64   `json:"amount"`
	Approval        string    `json:"approval"`
	ReferenceNumber string    `json:"referenceNumber"`
	Avs             string    `json:"avs"`
	ResponseMessage string    `json:"responseMessage"`
	ResponseCode    string    `json:"responseCode"`
	Gwid            int64     `json:"gwid"`
	Success         bool      `json:"success"`
	OrderID         int64     `json:"orderId"`
}

//LocalDataStore LocalDataStore
type LocalDataStore struct {
	StoreID       int64     `json:"storeId"`
	DataStoreName string    `json:"dataStoreName"`
	Reload        bool      `json:"reload"`
	ReloadDate    time.Time `json:"reloadDate"`
}

//Instances Instances
type Instances struct {
	InstanceName  string    `json:"instanceName"`
	ReloadDate    time.Time `json:"reloadDate"`
	StoreID       int64     `json:"storeId"`
	DataStoreName string    `json:"dataStoreName"`
}

//DataStoreWriteLock DataStoreWriteLock
type DataStoreWriteLock struct {
	ID                 int64     `json:"id"`
	DataStoreName      string    `json:"dataStoreName"`
	Locked             bool      `json:"locked"`
	LockedInstanceName string    `json:"lockedInstanceName"`
	LockedTime         time.Time `json:"lockedTime"`
	LockedByUser       string    `json:"lockedByUser"`
	StoreID            int64     `json:"storeId"`
}
