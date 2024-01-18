package parm

type Item struct {
IID							int		`json:"IID" gorm:"column:IID;not null"`
IName						string 	`json:"IName" gorm:"column:IName;size:40"`
IType						int		`json:"IType" gorm:"column:IType"`
ILevel						uint8	`json:"ILevel" gorm:"column:ILevel"`
IDHIT						int16	`json:"IDHIT" gorm:"column:IDHIT"`
IDDD						string	`json:"IDDD" gorm:"column:IDDD;size:10"`
IRHIT						int16	`json:"IRHIT" gorm:"column:IRHIT"`
IRDD						string	`json:"IRDD" gorm:"column:IRDD;size:50"`
IMHIT						int16	`json:"IMHIT" gorm:"column:IMHIT"`
IMDD						string	`json:"IMDD" gorm:"column:IMDD;size:10"`
IHPPlus						int16	`json:"IHPPlus" gorm:"column:IHPPlus"`
IMPPlus						int16	`json:"IMPPlus" gorm:"column:IMPPlus"`
ISTR						int16	`json:"ISTR" gorm:"column:ISTR"`
IDEX						int16	`json:"IDEX" gorm:"column:IDEX"`
IINT						int16	`json:"IINT" gorm:"column:IINT"`
IMaxStack					int		`json:"IMaxStack" gorm:"column:IMaxStack"`
IWeight						int16	`json:"IWeight" gorm:"column:IWeight"`
IUseType					int		`json:"IUseType" gorm:"column:IUseType"`
IUseNum						int		`json:"IUseNum" gorm:"column:IUseNum"`
IRecycle					int		`json:"IRecycle" gorm:"column:IRecycle"`
IHPRegen					uint8	`json:"IHPRegen" gorm:"column:IHPRegen"`
IMPRegen					uint8	`json:"IMPRegen" gorm:"column:IMPRegen"`
IAttackRate					uint8	`json:"IAttackRate" gorm:"column:IAttackRate"`
IMoveRate					uint8	`json:"IMoveRate" gorm:"column:IMoveRate"`
ICritical					uint8	`json:"ICritical" gorm:"column:ICritical"`
ITermOfValidity				int16	`json:"ITermOfValidity" gorm:"column:ITermOfValidity"`
ITermOfValidityMi			int16	`json:"ITermOfValidityMi" gorm:"column:ITermOfValidityMi"`
IDesc						string	`json:"IDesc" gorm:"column:IDesc;size:200"`
IStatus						uint8	`json:"IStatus" gorm:"column:IStatus"`
IFakeID						int		`json:"IFakeID" gorm:"column:IFakeID;not null"`
IFakeName					string	`json:"IFakeName" gorm:"column:IFakeName;not null;size:40"`
IUseMsg						string	`json:"IUseMsg" gorm:"column:IUseMsg;size:50"`
IRange						int16	`json:"IRange" gorm:"column:IRange;not null"`
IUseClass					uint8	`json:"IUseClass" gorm:"column:IUseClass;not null"`
IDropEffect					int		`json:"IDropEffect" gorm:"column:IDropEffect;not null"`
IUseLevel					int16	`json:"IUseLevel" gorm:"column:IUseLevel;not null"`
IUseEternal					uint8	`json:"IUseEternal" gorm:"column:IUseEternal;not null"`
IUseDelay					int		`json:"IUseDelay" gorm:"column:IUseDelay;not null"`
IUseInAttack				uint8	`json:"IUseInAttack" gorm:"column:IUseInAttack;not null"`
IIsEvent					bool	`json:"IIsEvent" gorm:"column:IIsEvent;not null"`
IIsIndict					bool	`json:"IIsIndict" gorm:"column:IIsIndict;not null"`
IAddWeight					int16	`json:"IAddWeight" gorm:"column:IAddWeight;not null"`
ISubType					int16	`json:"ISubType" gorm:"column:ISubType;not null"`
IIsCharge					bool	`json:"IIsCharge" gorm:"column:IIsCharge;not null"`
INationOp					int64	`json:"INationOp" gorm:"column:INationOp;not null"`
IPShopItemType				uint8	`json:"IPShopItemType" gorm:"column:IPShopItemType;not null"`
IQuestNo					int		`json:"IQuestNo" gorm:"column:IQuestNo;not null"`
IIsTest						bool	`json:"IIsTest" gorm:"column:IIsTest;not null"`
IQuestNeedCnt				uint8	`json:"IQuestNeedCnt" gorm:"column:IQuestNeedCnt;not null"`
IContentsLv					uint8	`json:"IContentsLv" gorm:"column:IContentsLv;not null"`
IIsConfirm					bool	`json:"IIsConfirm" gorm:"column:IIsConfirm;not null"`
IIsSealable					bool	`json:"IIsSealable" gorm:"column:IIsSealable;not null"`
IAddDDWhenCritical			int16	`json:"IAddDDWhenCritical" gorm:"column:IAddDDWhenCritical;not null"`
MSealRemovalNeedCnt			uint8	`json:"mSealRemovalNeedCnt" gorm:"column:mSealRemovalNeedCnt;not null"`
MIsPracticalPeriod			bool	`json:"mIsPracticalPeriod" gorm:"column:mIsPracticalPeriod;not null"`
MIsReceiveTown				bool	`json:"mIsReceiveTown" gorm:"column:mIsReceiveTown;not null"`
IIsReinforceDestroy			bool	`json:"IIsReinforceDestroy" gorm:"column:IIsReinforceDestroy;not null"`
IAddPotionRestore			int16	`json:"IAddPotionRestore" gorm:"column:IAddPotionRestore;not null"`
IAddMaxHpWhenTransform		int16	`json:"IAddMaxHpWhenTransform" gorm:"column:IAddMaxHpWhenTransform;not null"`
IAddMaxMpWhenTransform		int16	`json:"IAddMaxMpWhenTransform" gorm:"column:IAddMaxMpWhenTransform;not null"`
IAddAttackRateWhenTransform	int16	`json:"IAddAttackRateWhenTransform" gorm:"column:IAddAttackRateWhenTransform;not null"`
IAddMoveRateWhenTransform	int16	`json:"IAddMoveRateWhenTransform" gorm:"column:IAddMoveRateWhenTransform;not null"`
ISupportType				uint8	`json:"ISupportType" gorm:"column:ISupportType;not null"`
ITermOfValidityLv			int16	`json:"ITermOfValidityLv" gorm:"column:ITermOfValidityLv;not null"`
MIsUseableUTGWSvr			bool	`json:"mIsUseableUTGWSvr" gorm:"column:mIsUseableUTGWSvr;not null"`
IAddShortAttackRange		int16	`json:"IAddShortAttackRange" gorm:"column:IAddShortAttackRange;not null"`
IAddLongAttackRange			int16	`json:"IAddLongAttackRange" gorm:"column:IAddLongAttackRange;not null"`
IWeaponPoisonType			int16	`json:"IWeaponPoisonType" gorm:"column:IWeaponPoisonType;not null"`
IDPV						int16	`json:"IDPV" gorm:"column:IDPV;not null"`
IMPV						int16	`json:"IMPV" gorm:"column:IMPV;not null"`
IRPV						int16	`json:"IRPV" gorm:"column:IRPV;not null"`
IDDV						int16	`json:"IDDV" gorm:"column:IDDV;not null"`
IMDV						int16	`json:"IMDV" gorm:"column:IMDV;not null"`
IRDV						int16	`json:"IRDV" gorm:"column:IRDV;not null"`
IHDPV						int16	`json:"IHDPV" gorm:"column:IHDPV;not null"`
IHMPV						int16	`json:"IHMPV" gorm:"column:IHMPV;not null"`
IHRPV						int16	`json:"IHRPV" gorm:"column:IHRPV;not null"`
IHDDV						int16	`json:"IHDDV" gorm:"column:IHDDV;not null"`
IHMDV						int16	`json:"IHMDV" gorm:"column:IHMDV;not null"`
IHRDV						int16	`json:"IHRDV" gorm:"column:IHRDV;not null"`
ISubDDWhenCritical			int16	`json:"ISubDDWhenCritical" gorm:"column:ISubDDWhenCritical;not null"`
IGetItemFeedback			int16	`json:"IGetItemFeedback" gorm:"column:IGetItemFeedback;not null"`
IEnemySubCriticalHit		int16	`json:"IEnemySubCriticalHit" gorm:"column:IEnemySubCriticalHit;not null"`
IIsPartyDrop				bool	`json:"IIsPartyDrop" gorm:"column:IIsPartyDrop;not null"`
IMaxBeadHoleCount			uint8	`json:"IMaxBeadHoleCount" gorm:"column:IMaxBeadHoleCount;not null"`
ISubTypeOption				int		`json:"ISubTypeOption" gorm:"column:ISubTypeOption;not null"`
MIsDeleteArenaSvr			bool	`json:"mIsDeleteArenaSvr" gorm:"column:mIsDeleteArenaSvr;not null"`
}

func (Item) TableName() string {
	return "DT_Item"
}