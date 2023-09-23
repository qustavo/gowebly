package validators

import (
	"errors"

	"github.com/gowebly/gowebly/internal/config"
	"github.com/gowebly/gowebly/internal/constants"
)

// Validate validates the given config.
func Validate(cfg *config.Config) error {
	// Check, if 'backend' block is set.
	if cfg.Backend == nil {
		return errors.New(constants.ErrorValidateConfigBackendBlockNotFound)
	}

	// Check, if 'backend' block is set.
	if cfg.Backend != nil {
		// Check, if 'name' option in the 'backend' block is not set.
		if cfg.Backend.Name == "" {
			return errors.New(constants.ErrorValidateConfigBackendNameNotFound)
		}

		// Check, if 'port' option in the 'backend' block is not set.
		if cfg.Backend.Port == 0 {
			return errors.New(constants.ErrorValidateConfigBackendPortNotFound)
		}

		// Check, if 'timeout' option in the 'backend' block is not set.
		if cfg.Backend.Timeout == nil {
			return errors.New(constants.ErrorValidateConfigBackendTimeoutNotFound)
		}
	}

	// Check, if the 'frontend' block is not set.
	if cfg.Frontend == nil {
		return errors.New(constants.ErrorValidateConfigFrontendBlockNotFound)
	}

	// Check, if the 'frontend' block is set.
	if cfg.Frontend != nil {
		// Check, if 'htmx' option in the 'frontend' block is not set.
		if cfg.Frontend.HTMX == "" {
			return errors.New(constants.ErrorValidateConfigFrontendHTMXNotFound)
		}

		// Check, if 'hyperscript' option in the 'frontend' block is not set.
		if cfg.Frontend.Hyperscript == "" {
			return errors.New(constants.ErrorValidateConfigFrontendHyperscriptNotFound)
		}
	}

	return nil
}