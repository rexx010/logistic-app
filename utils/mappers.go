package utils

import (
	"logisticApp/data/models"
	"logisticApp/dtos/responses"

	"github.com/google/uuid"
)

func ToUserResponse(user models.User) responses.UserResponse {
	resp := responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	if user.BusinessProfile != nil {
		bp := ToBusinessProfileResponse(*user.BusinessProfile)
		resp.BusinessProfile = &bp
	}
	if user.RiderProfile != nil {
		rp := ToRiderProfileResponse(*user.RiderProfile)
		resp.RiderProfile = &rp
	}
	return resp
}

func ToDeliveryResponse(delivery models.Delivery) responses.DeliveryResponse {
	resp := responses.DeliveryResponse{
		ID:             delivery.ID,
		BusinessID:     delivery.BusinessID,
		RiderID:        delivery.RiderID,
		PickupAddress:  delivery.PickupAddress,
		DropoffAddress: delivery.DropoffAddress,
		DeliveryFee:    delivery.DeliveryFee,
		Status:         delivery.Status,
		ProofImageURL:  delivery.ProofImageURL,
		CreatedAt:      delivery.CreatedAt,
		UpdatedAt:      delivery.UpdatedAt,
	}
	if delivery.Rider != nil {
		r := ToRiderPublicResponse(*delivery.Rider, delivery.Rider.User)
		resp.Rider = &r
	}
	if delivery.Payment != nil {
		p := ToPaymentResponse(*delivery.Payment)
		resp.Payment = &p
	}
	if delivery.Invoice != nil {
		inv := ToInvoiceResponse(*delivery.Invoice)
		resp.Invoice = &inv
	}
	if delivery.Commission != nil {
		c := ToCommissionResponse(*delivery.Commission)
		resp.Commission = &c
	}
	return resp
}

func ToPaymentResponse(payment models.Payment) responses.PaymentResponse {
	return responses.PaymentResponse{
		ID:               payment.ID,
		DeliveryID:       payment.DeliveryID,
		Amount:           payment.Amount,
		Status:           payment.Status,
		PaymentReference: payment.PaymentReference,
		PaymentMethod:    payment.PaymentMethod,
		CreatedAt:        payment.CreatedAt,
		UpdatedAt:        payment.UpdatedAt,
	}
}

func ToNotificationResponse(notification models.Notification) responses.NotificationResponse {
	return responses.NotificationResponse{
		ID:          notification.ID,
		UserID:      notification.UserID,
		Type:        notification.Type,
		Message:     notification.Message,
		TriggeredBy: notification.TriggeredBy,
		Status:      notification.Status,
		SentAt:      notification.SentAt,
		CreatedAt:   notification.CreatedAt,
	}
}

func ToRiderProfileResponse(rp models.RiderProfile) responses.RiderProfileResponse {
	id, _ := uuid.Parse(rp.UserID)
	return responses.RiderProfileResponse{
		ID:              rp.ID,
		UserID:          id,
		VehicleType:     rp.VehicleType,
		LicenseNumber:   rp.LicenseNumber,
		EarningsBalance: rp.EarningsBalance,
		Status:          rp.Status,
		CreatedAt:       rp.CreatedAt,
		UpdatedAt:       rp.UpdatedAt,
	}
}

func ToRiderPublicResponse(rp models.RiderProfile, u models.User) responses.RiderPublicResponse {
	id, _ := uuid.Parse(rp.UserID)
	return responses.RiderPublicResponse{
		ID:            rp.ID,
		UserID:        id,
		Name:          u.Name,
		Phone:         u.Phone,
		VehicleType:   rp.VehicleType,
		LicenseNumber: rp.LicenseNumber,
		Status:        rp.Status,
	}
}

func ToInvoiceResponse(invoice models.Invoice) responses.InvoiceResponse {
	return responses.InvoiceResponse{
		ID:         invoice.ID,
		DeliveryID: invoice.DeliveryID,
		Amount:     invoice.Amount,
		Status:     invoice.Status,
		PDFURL:     invoice.PDFURL,
		CreatedAt:  invoice.CreatedAt,
		UpdatedAt:  invoice.UpdatedAt,
	}
}

func ToCommissionResponse(commission models.Commission) responses.CommissionResponse {
	return responses.CommissionResponse{
		ID:          commission.ID,
		DeliveryID:  commission.DeliveryID,
		PlatformFee: commission.PlatformFee,
		RiderFee:    commission.RiderFee,
		CreatedAt:   commission.CreatedAt,
		UpdatedAt:   commission.UpdatedAt,
	}
}

func ToBusinessProfileResponse(bp models.BusinessProfile) responses.BusinessProfileResponse {
	id, _ := uuid.Parse(bp.UserID)
	return responses.BusinessProfileResponse{
		ID:               bp.ID,
		UserID:           id,
		BusinessName:     bp.BusinessName,
		BusinessAddress:  bp.BusinessAddress,
		SubscriptionPlan: bp.SubscriptionPlan,
		Status:           bp.Status,
		CreatedAt:        bp.CreatedAt,
		UpdatedAt:        bp.UpdatedAt,
	}
}
