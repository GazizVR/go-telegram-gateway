package telegramgateway

import "context"

const methodSendVerificationMessage = "sendVerificationMessage"

func (c *Client) SendVerificationMessage(
	ctx context.Context,
	req SendVerificationMessageRequest,
) (*RequestStatus, error) {
	var status RequestStatus
	if err := c.call(ctx, methodSendVerificationMessage, req, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

const methodCheckSendAbility = "checkSendAbility"

func (c *Client) CheckSendAbility(
	ctx context.Context,
	req CheckSendAbilityRequest,
) (*RequestStatus, error) {
	var status RequestStatus
	if err := c.call(ctx, methodCheckSendAbility, req, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

const methodCheckVerificationStatus = "checkVerificationStatus"

func (c *Client) CheckVerificationStatus(
	ctx context.Context,
	req CheckVerificationStatusRequest,
) (*RequestStatus, error) {
	var status RequestStatus
	if err := c.call(ctx, methodCheckVerificationStatus, req, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

const methodRevokeVerificationMessage = "revokeVerificationMessage"

func (c *Client) RevokeVerificationMessage(
	ctx context.Context,
	req RevokeVerificationMessageRequest,
) (*bool, error) {
	var isRevoked bool
	if err := c.call(ctx, methodRevokeVerificationMessage, req, &isRevoked); err != nil {
		return nil, err
	}
	return &isRevoked, nil
}
