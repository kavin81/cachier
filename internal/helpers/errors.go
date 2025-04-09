package helpers

func wrapError(err error, msg string) error {
	logger := Log

	if err == nil {
		return nil
	}
	logger.Error(msg, err)

	return err
}

