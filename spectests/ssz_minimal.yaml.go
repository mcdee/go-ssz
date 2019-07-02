// Code generated by yaml_to_go. DO NOT EDIT.
// source: ssz_minimal_one.yaml

package autogenerated

type MinimalFork struct {
	PreviousVersion []byte `json:"previous_version" ssz:"size=4"`
	CurrentVersion  []byte `json:"current_version" ssz:"size=4"`
	Epoch           uint64 `json:"epoch"`
}

type MinimalCheckpoint struct {
	Epoch uint64 `json:"epoch"`
	Root  []byte `json:"root" ssz:"size=32"`
}

type MinimalValidator struct {
	Pubkey                     []byte `json:"pubkey" ssz:"size=48"`
	WithdrawalCredentials      []byte `json:"withdrawal_credentials" ssz:"size=32"`
	EffectiveBalance           uint64 `json:"effective_balance"`
	Slashed                    bool   `json:"slashed"`
	ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch"`
	ActivationEpoch            uint64 `json:"activation_epoch"`
	ExitEpoch                  uint64 `json:"exit_epoch"`
	WithdrawableEpoch          uint64 `json:"withdrawable_epoch"`
}

type MinimalCrosslink struct {
	Shard      uint64 `json:"shard"`
	ParentRoot []byte `json:"parent_root" ssz:"size=32"`
	StartEpoch uint64 `json:"start_epoch"`
	EndEpoch   uint64 `json:"end_epoch"`
	DataRoot   []byte `json:"data_root" ssz:"size=32"`
}

type MinimalAttestationData struct {
	BeaconBlockRoot []byte            `json:"beacon_block_root" ssz:"size=32"`
	Source          MinimalCheckpoint `json:"source"`
	Target          MinimalCheckpoint `json:"target"`
	Crosslink       MinimalCrosslink  `json:"crosslink"`
}

type MinimalAttestationAndCustodyBit struct {
	Data       MinimalAttestationData `json:"data"`
	CustodyBit byte                   `json:"custody_bit"`
}

type MinimalIndexedAttestation struct {
	CustodyBit0Indices []uint64               `json:"custody_bit_0_indices"`
	CustodyBit1Indices []uint64               `json:"custody_bit_1_indices"`
	Data               MinimalAttestationData `json:"data"`
	Signature          []byte                 `json:"signature" ssz:"size=96"`
}

type MinimalPendingAttestation struct {
	AggregationBitfield []byte                 `json:"aggregation_bitfield"`
	Data                MinimalAttestationData `json:"data"`
	InclusionDelay      uint64                 `json:"inclusion_delay"`
	ProposerIndex       uint64                 `json:"proposer_index"`
}

type MinimalEth1Data struct {
	DepositRoot  []byte `json:"deposit_root" ssz:"size=32"`
	DepositCount uint64 `json:"deposit_count"`
	BlockHash    []byte `json:"block_hash" ssz:"size=32"`
}

type MinimalHistoricalBatch struct {
	BlockRoots [][]byte `json:"block_roots" ssz:"size=64,32"`
	StateRoots [][]byte `json:"state_roots" ssz:"size=64,32"`
}

type MinimalDepositData struct {
	Pubkey                []byte `json:"pubkey" ssz:"size=48"`
	WithdrawalCredentials []byte `json:"withdrawal_credentials" ssz:"size=32"`
	Amount                uint64 `json:"amount"`
	Signature             []byte `json:"signature" ssz:"size=96"`
}

type MinimalCompactCommittee struct {
	Pubkeys           []byte   `json:"pubkeys" ssz:"size=48"`
	CompactValidators []uint64 `json:"compact_validators"`
}

// TODO Header
type MinimalBlockHeader struct {
	Slot       uint64 `json:"slot"`
	ParentRoot []byte `json:"parent_root" ssz:"size=32"`
	StateRoot  []byte `json:"state_root" ssz:"size=32"`
	BodyRoot   []byte `json:"body_root" ssz:"size=32"`
	Signature  []byte `json:"signature" ssz:"size=96"`
}

type MinimalProposerSlashing struct {
	ProposerIndex uint64             `json:"proposer_index"`
	Header1       MinimalBlockHeader `json:"header_1"`
	Header2       MinimalBlockHeader `json:"header_2"`
}

type MinimalAttesterSlashing struct {
	Attestation1 MinimalIndexedAttestation `json:"attestation_1"`
	Attestation2 MinimalIndexedAttestation `json:"attestation_2"`
}

type MinimalAttestation struct {
	AggregationBits []byte                 `json:"aggregation_bitfield" ssz:"size=4096"`
	Data            MinimalAttestationData `json:"data"`
	CustodyBits     []byte                 `json:"custody_bitfield" ssz:"size=4096"`
	Signature       []byte                 `json:"signature" ssz:"size=96"`
}

type MinimalDeposit struct {
	Proof [][]byte           `json:"proof" ssz:"size=33,32"`
	Data  MinimalDepositData `json:"data"`
}

type MinimalVoluntaryExit struct {
	Epoch          uint64 `json:"epoch"`
	ValidatorIndex uint64 `json:"validator_index"`
	Signature      []byte `json:"signature" ssz:"size=96"`
}

type MinimalTransfer struct {
	Sender    uint64 `json:"sender"`
	Recipient uint64 `json:"recipient"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
	Slot      uint64 `json:"slot"`
	Pubkey    []byte `json:"pubkey" ssz:"size=48"`
	Signature []byte `json:"signature" ssz:"size=96"`
}

type MinimalBlockBody struct {
	RandaoReveal      []byte                    `json:"randao_reveal" ssz:"size=96"`
	Eth1Data          MinimalEth1Data           `json:"eth1_data"`
	Graffiti          []byte                    `json:"graffiti" ssz:"size=32"`
	ProposerSlashings []MinimalProposerSlashing `json:"proposer_slashings"`
	AttesterSlashings []MinimalAttesterSlashing `json:"attester_slashings"`
	Attestations      []MinimalAttestation      `json:"attestations"`
	Deposits          []MinimalDeposit          `json:"deposits"`
	VoluntaryExits    []MinimalVoluntaryExit    `json:"voluntary_exits"`
	Transfers         []MinimalTransfer         `json:"transfers"`
}

type MinimalBlock struct {
	Slot       uint64           `json:"slot"`
	ParentRoot []byte           `json:"parent_root" ssz:"size=32"`
	StateRoot  []byte           `json:"state_root" ssz:"size=32"`
	Body       MinimalBlockBody `json:"body"`
	Signature  []byte           `json:"signature" ssz:"size=96"`
}

type MinimalBeaconState struct {
	Slot                   uint64             `json:"slot"`
	GenesisTime            uint64             `json:"genesis_time"`
	Fork                   MinimalFork        `json:"fork"`
	LatestBlockHeader      MinimalBlockHeader `json:"latest_block_header"`
	BlockRoots             [][]byte           `json:"latest_block_roots" ssz:"size=64,32"`
	StateRoots             [][]byte           `json:"latest_state_roots" ssz:"size=64,32"`
	HistoricalRoots        [][]byte           `json:"historical_roots" ssz:"size=?,32"`
	Eth1Data               MinimalEth1Data    `json:"latest_eth1_data"`
	Eth1DataVotes          []MinimalEth1Data  `json:"eth1_data_votes"`
	Eth1DepositIndex       uint64             `json:"deposit_index"`
	Validators             []MinimalValidator `json:"validator_registry"`
	Balances               []uint64           `json:"balances"`
	StartShard             uint64             `json:"latest_start_shard"`
	RandaoMixes            [][]byte           `json:"latest_randao_mixes" ssz:"size=64,32"`
	ActiveIndexRoots       [][]byte           `json:"latest_active_index_roots" ssz:"size=64,32"`
	CompactCommitteesRoots [][]byte           `json:"compact_committees_roots" ssz:"size=64,32"`
	Slashings              []uint64           `json:"latest_slashed_balances" ssz:"size=64"`

	PreviousEpochAttestations []MinimalPendingAttestation `json:"previous_epoch_attestations"`
	CurrentEpochAttestations  []MinimalPendingAttestation `json:"current_epoch_attestations"`
	PreviousCrosslinks        []MinimalCrosslink          `json:"previous_crosslinks" ssz:"size=8"`
	CurrentCrosslinks         []MinimalCrosslink          `json:"current_crosslinks" ssz:"size=8"`
	JustificationBits         []byte                      `json:"justification_bitfield"`

	PreviousJustifiedCheckpoint MinimalCheckpoint `json:"previous_justified_epoch"`
	CurrentJustifiedCheckpoint  MinimalCheckpoint `json:"current_justified_epoch"`
	FinalizedCheckpoint         MinimalCheckpoint `json:"finalized_checkpoint"`
}
type SszMinimalTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Attestation struct {
			Value       MinimalAttestation `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"Attestation,omitempty"`
		AttestationData struct {
			Value      MinimalAttestationData `json:"value"`
			Serialized []byte                 `json:"serialized"`
			Root       []byte                 `json:"root" ssz:"size=32"`
		} `json:"AttestationData,omitempty"`
		AttestationDataAndCustodyBit struct {
			Value      MinimalAttestationAndCustodyBit `json:"value"`
			Serialized []byte                          `json:"serialized"`
			Root       []byte                          `json:"root" ssz:"size=32"`
		} `json:"AttestationDataAndCustodyBit,omitempty"`
		AttesterSlashing struct {
			Value      MinimalAttesterSlashing `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"AttesterSlashing,omitempty"`
		BeaconBlock struct {
			Value       MinimalBlock `json:"value"`
			Serialized  []byte       `json:"serialized"`
			Root        []byte       `json:"root" ssz:"size=32"`
			SigningRoot []byte       `json:"signing_root" ssz:"size=32"`
		} `json:"BeaconBlock,omitempty"`
		BeaconBlockBody struct {
			Value      MinimalBlockBody `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"BeaconBlockBody,omitempty"`
		BeaconBlockHeader struct {
			Value       MinimalBlockHeader `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"BeaconBlockHeader,omitempty"`
		BeaconState struct {
			Value      MinimalBeaconState `json:"value"`
			Serialized []byte             `json:"serialized"`
			Root       []byte             `json:"root" ssz:"size=32"`
		} `json:"BeaconState,omitempty"`
		Crosslink struct {
			Value      MinimalCrosslink `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"Crosslink,omitempty"`
		Deposit struct {
			Value      MinimalDeposit `json:"value"`
			Serialized []byte         `json:"serialized"`
			Root       []byte         `json:"root" ssz:"size=32"`
		} `json:"Deposit,omitempty"`
		DepositData struct {
			Value       MinimalDepositData `json:"value"`
			Serialized  []byte             `json:"serialized"`
			Root        []byte             `json:"root" ssz:"size=32"`
			SigningRoot []byte             `json:"signing_root" ssz:"size=32"`
		} `json:"DepositData,omitempty"`
		Eth1Data struct {
			Value      MinimalEth1Data `json:"value"`
			Serialized []byte          `json:"serialized"`
			Root       []byte          `json:"root" ssz:"size=32"`
		} `json:"Eth1Data,omitempty"`
		Fork struct {
			Value      MinimalFork `json:"value"`
			Serialized []byte      `json:"serialized"`
			Root       []byte      `json:"root" ssz:"size=32"`
		} `json:"Fork,omitempty"`
		HistoricalBatch struct {
			Value      MinimalHistoricalBatch `json:"value"`
			Serialized []byte                 `json:"serialized"`
			Root       []byte                 `json:"root" ssz:"size=32"`
		} `json:"HistoricalBatch,omitempty"`
		IndexedAttestation struct {
			Value       MinimalIndexedAttestation `json:"value"`
			Serialized  []byte                    `json:"serialized"`
			Root        []byte                    `json:"root" ssz:"size=32"`
			SigningRoot []byte                    `json:"signing_root" ssz:"size=32"`
		} `json:"IndexedAttestation,omitempty"`
		PendingAttestation struct {
			Value      MinimalPendingAttestation `json:"value"`
			Serialized []byte                    `json:"serialized"`
			Root       []byte                    `json:"root" ssz:"size=32"`
		} `json:"PendingAttestation,omitempty"`
		ProposerSlashing struct {
			Value      MinimalProposerSlashing `json:"value"`
			Serialized []byte                  `json:"serialized"`
			Root       []byte                  `json:"root" ssz:"size=32"`
		} `json:"ProposerSlashing,omitempty"`
		Transfer struct {
			Value       MinimalTransfer `json:"value"`
			Serialized  []byte          `json:"serialized"`
			Root        []byte          `json:"root" ssz:"size=32"`
			SigningRoot []byte          `json:"signing_root" ssz:"size=32"`
		} `json:"Transfer,omitempty"`
		Validator struct {
			Value      MinimalValidator `json:"value"`
			Serialized []byte           `json:"serialized"`
			Root       []byte           `json:"root" ssz:"size=32"`
		} `json:"Validator,omitempty"`
		VoluntaryExit struct {
			Value       MinimalVoluntaryExit `json:"value"`
			Serialized  []byte               `json:"serialized"`
			Root        []byte               `json:"root" ssz:"size=32"`
			SigningRoot []byte               `json:"signing_root" ssz:"size=32"`
		} `json:"VoluntaryExit,omitempty"`
	} `json:"test_cases"`
}
