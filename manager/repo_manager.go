package manager

import (
	"bank-api/repository"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepository
	GetMerchantRepo() repository.MerchantRepository
	
}

type repoManager struct {
	
	usrRepo      repository.UserRepository
	mctRepo 	 repository.MerchantRepository
}

var onceLoadUserRepo sync.Once
var onceLoadMerchantRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepository {
	onceLoadUserRepo.Do(func() {
		rm.usrRepo = repository.NewUserRepository()
	})
	return rm.usrRepo
}

func (rm *repoManager) GetMerchantRepo() repository.MerchantRepository {
	onceLoadMerchantRepo.Do(func() {
		rm.mctRepo = repository.NewMerchantRepository()
	})
	return rm.mctRepo
}

func NewRepoManager() RepoManager {
	return &repoManager{}
}