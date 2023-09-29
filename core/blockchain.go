package core

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	store     Storage
	headers   []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemStore(),
	}

	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	err := bc.validator.ValidateBlock(b)
	if err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) getHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height (%d) is too high", height)
	}

	return bc.headers[height], nil
}

func (bc *Blockchain) hasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	logrus.WithFields(logrus.Fields{
		"Height": b.Height,
		"Hash":   b.Hash(BlockHasher{}),
	}).Info("Adding new block")
	bc.headers = append(bc.headers, b.Header)

	return bc.store.Put(b)
}
