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
