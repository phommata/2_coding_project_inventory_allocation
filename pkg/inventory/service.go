package inventory

type InventoryService interface {
	VoidSale(TransactionData *entities.TransactionData) (*entities.ProtobaseTransactionBatch, error)
}

type inventoryService struct {
	logger logger.Logger
}

func NewInventoryService(logger logger.Logger, client Client) *inventoryService {
	return &inventoryService{
		logger: logger,
	}
}

func (svc *inventoryService) VoidSale(TransactionData *entities.TransactionData) (*entities.ProtobaseTransactionBatch, error) {
	if GatewayConfig.URL != "" {
		svc.client.URL = GatewayConfig.URL
	}

	//ProtobaseTransactionBatch, err := svc.client.VoidSale(GatewayConfig, TransactionData)

	return ProtobaseTransactionBatch, err
}
