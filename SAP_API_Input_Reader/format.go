package sap_api_input_reader

import (
	"sap-api-integrations-sales-order-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeaderItem() *requests.HeaderItem {
	data := sdc.SalesOrder
	results := make([]requests.Item, 0, len(data.SalesOrderItem))

	header := sdc.ConvertToHeader()

	for i := range data.SalesOrderItem {
		results = append(results, *sdc.ConvertToItem(i))
	}

	return &requests.HeaderItem{
		Header: *header,
		ToItem: requests.ToItem{
			Results: results,
		},
	}
}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.SalesOrder
	return &requests.Header{
		SalesOrder:                data.SalesOrder,
		SalesOrderType:            data.SalesOrderType,
		SalesOrganization:         data.SalesOrganization,
		DistributionChannel:       data.DistributionChannel,
		OrganizationDivision:      data.OrganizationDivision,
		SalesGroup:                data.SalesGroup,
		SalesOffice:               data.SalesOffice,
		SalesDistrict:             data.SalesDistrict,
		SoldToParty:               data.SoldToParty,
		CreationDate:              data.CreationDate,
		LastChangeDate:            data.LastChangeDate,
		ExternalDocumentID:        data.ExternalDocumentID,
		LastChangeDateTime:        data.LastChangeDateTime,
		PurchaseOrderByCustomer:   data.PurchaseOrderByCustomer,
		CustomerPurchaseOrderDate: data.CustomerPurchaseOrderDate,
		SalesOrderDate:            data.SalesOrderDate,
		TotalNetAmount:            data.TotalNetAmount,
		//		OverallDeliveryStatus:          data.OverallDeliveryStatus,
		//		TotalBlockStatus:               data.TotalBlockStatus,
		//		OverallOrdReltdBillgStatus:     data.OverallOrdReltdBillgStatus,
		//		OverallSDDocReferenceStatus:    data.OverallSDDocReferenceStatus,
		TransactionCurrency: data.TransactionCurrency,
		SDDocumentReason:    data.SDDocumentReason,
		PricingDate:         data.PricingDate,
		//		PriceDetnExchangeRate:          data.PriceDetnExchangeRate,
		RequestedDeliveryDate:          data.RequestedDeliveryDate,
		ShippingCondition:              data.ShippingCondition,
		CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
		ShippingType:                   data.ShippingType,
		HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
		DeliveryBlockReason:            data.DeliveryBlockReason,
		IncotermsClassification:        data.IncotermsClassification,
		CustomerPriceGroup:             data.CustomerPriceGroup,
		PriceListType:                  data.PriceListType,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		ReferenceSDDocument:            data.ReferenceSDDocument,
		ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
		CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
		AccountingExchangeRate:         data.AccountingExchangeRate,
		CustomerGroup:                  data.CustomerGroup,
		AdditionalCustomerGroup1:       data.AdditionalCustomerGroup1,
		AdditionalCustomerGroup2:       data.AdditionalCustomerGroup2,
		AdditionalCustomerGroup3:       data.AdditionalCustomerGroup3,
		AdditionalCustomerGroup4:       data.AdditionalCustomerGroup4,
		AdditionalCustomerGroup5:       data.AdditionalCustomerGroup5,
		CustomerTaxClassification1:     data.CustomerTaxClassification1,
		TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
		//		BillingDocumentDate:            data.BillingDocumentDate,
	}
}

func (sdc *SDC) ConvertToItem(num int) *requests.Item {
	dataSalesOrder := sdc.SalesOrder
	data := sdc.SalesOrder.SalesOrderItem[num]
	return &requests.Item{
		SalesOrder:              dataSalesOrder.SalesOrder,
		SalesOrderItem:          data.SalesOrderItem,
		SalesOrderItemCategory:  data.SalesOrderItemCategory,
		SalesOrderItemText:      data.SalesOrderItemText,
		PurchaseOrderByCustomer: data.PurchaseOrderByCustomer,
		Material:                data.Material,
		MaterialByCustomer:      data.MaterialByCustomer,
		PricingDate:             data.PricingDate,
		//		BillingPlan:                 data.BillingPlan,
		RequestedQuantity:     data.RequestedQuantity,
		RequestedQuantityUnit: data.RequestedQuantityUnit,
		//		OrderQuantityUnit:           data.OrderQuantityUnit,
		//		ConfdDelivQtyInOrderQtyUnit: data.ConfdDelivQtyInOrderQtyUnit,
		ItemGrossWeight:      data.ItemGrossWeight,
		ItemNetWeight:        data.ItemNetWeight,
		ItemWeightUnit:       data.ItemWeightUnit,
		ItemVolume:           data.ItemVolume,
		ItemVolumeUnit:       data.ItemVolumeUnit,
		TransactionCurrency:  data.TransactionCurrency,
		NetAmount:            data.NetAmount,
		MaterialGroup:        data.MaterialGroup,
		MaterialPricingGroup: data.MaterialPricingGroup,
		//		BillingDocumentDate:         data.BillingDocumentDate,
		Batch:                   data.Batch,
		ProductionPlant:         data.ProductionPlant,
		StorageLocation:         data.StorageLocation,
		DeliveryGroup:           data.DeliveryGroup,
		ShippingPoint:           data.ShippingPoint,
		ShippingType:            data.ShippingType,
		DeliveryPriority:        data.DeliveryPriority,
		IncotermsClassification: data.IncotermsClassification,
		//		TaxAmount:                   data.TaxAmount,
		ProductTaxClassification1:  data.ProductTaxClassification1,
		MatlAccountAssignmentGroup: data.MatlAccountAssignmentGroup,
		//		CostAmount:                  data.CostAmount,
		CustomerPaymentTerms:      data.CustomerPaymentTerms,
		CustomerGroup:             data.CustomerGroup,
		SalesDocumentRjcnReason:   data.SalesDocumentRjcnReason,
		ItemBillingBlockReason:    data.ItemBillingBlockReason,
		WBSElement:                data.WBSElement,
		ProfitCenter:              data.ProfitCenter,
		AccountingExchangeRate:    data.AccountingExchangeRate,
		ReferenceSDDocument:       data.ReferenceSDDocument,
		ReferenceSDDocumentItem:   data.ReferenceSDDocumentItem,
		SDProcessStatus:           data.SDProcessStatus,
		DeliveryStatus:            data.DeliveryStatus,
		OrderRelatedBillingStatus: data.OrderRelatedBillingStatus,
	}
}
