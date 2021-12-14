package migrate

import (
	internalmongodb "github.com/koderover/zadig/pkg/cli/upgradeassistant/internal/repository/mongodb"
	"github.com/koderover/zadig/pkg/cli/upgradeassistant/internal/upgradepath"
	"github.com/koderover/zadig/pkg/tool/log"
)

func init() {
	upgradepath.AddHandler(upgradepath.V170, upgradepath.V180, V17xToV180)
	upgradepath.AddHandler(upgradepath.V171, upgradepath.V180, V17xToV180)
}

func V17xToV180() error {
	log.Info("Migrating data from 1.7.0 to 1.8.0")

	log.Info("Start to patchProductRegistryID")
	err := patchProductRegistryID()
	if err != nil {
		log.Errorf("Failed to patchProductRegistryID, err: %s", err)
		return err
	}

	return nil
}

func patchProductRegistryID() error {
	// get all products
	products, err := internalmongodb.NewProductColl().List(&internalmongodb.ProductListOptions{})
	if err != nil {
		log.Errorf("fail to list products, err: %s", err)
		return err
	}
	registry, err := internalmongodb.NewRegistryNamespaceColl().Find(&internalmongodb.FindRegOps{IsDefault: true})
	if err != nil {
		log.Errorf("fail to find default registry, err: %s", err)
		return err
	}
	// change type to readable string
	for _, v := range products {
		if len(v.RegistryID) == 0 {
			v.RegistryID = registry.ID.Hex()
		}
	}
	err = internalmongodb.NewProductColl().UpdateAllRegistry(products)
	if err != nil {
		log.Errorf("fail to update products, err: %s", err)
		return err
	}
	return nil
}
