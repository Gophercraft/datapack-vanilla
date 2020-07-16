package main

import (
	"fmt"
	"io"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/superp00t/etc/yo"
	"github.com/superp00t/gophercraft/datapack"
	"github.com/superp00t/gophercraft/datapack/text"
	"github.com/superp00t/gophercraft/econ"
	"github.com/superp00t/gophercraft/i18n"
	"github.com/superp00t/gophercraft/packet/update"
	"github.com/superp00t/gophercraft/worldserver/wdb"
	"xorm.io/xorm"
)

type NPCText struct {
	ID      uint32  `xorm:"'id'"`
	Text0_0 string  `xorm:"'text0_0'"`
	Text0_1 string  `xorm:"'text0_1'"`
	Lang0   uint32  `xorm:"'lang0'"`
	Prob0   float32 `xorm:"'prob0'"`
	Em0_0   uint32  `xorm:"'em0_0'"`
	Em0_1   uint32  `xorm:"'em0_1'"`
	Em0_2   uint32  `xorm:"'em0_2'"`
	Em0_3   uint32  `xorm:"'em0_3'"`
	Em0_4   uint32  `xorm:"'em0_4'"`
	Em0_5   uint32  `xorm:"'em0_5'"`
	Text1_0 string  `xorm:"'text1_0'"`
	Text1_1 string  `xorm:"'text1_1'"`
	Lang1   uint32  `xorm:"'lang1'"`
	Prob1   float32 `xorm:"'prob1'"`
	Em1_0   uint32  `xorm:"'em1_0'"`
	Em1_1   uint32  `xorm:"'em1_1'"`
	Em1_2   uint32  `xorm:"'em1_2'"`
	Em1_3   uint32  `xorm:"'em1_3'"`
	Em1_4   uint32  `xorm:"'em1_4'"`
	Em1_5   uint32  `xorm:"'em1_5'"`
	Text2_0 string  `xorm:"'text2_0'"`
	Text2_1 string  `xorm:"'text2_1'"`
	Lang2   uint32  `xorm:"'lang2'"`
	Prob2   float32 `xorm:"'prob2'"`
	Em2_0   uint32  `xorm:"'em2_0'"`
	Em2_1   uint32  `xorm:"'em2_1'"`
	Em2_2   uint32  `xorm:"'em2_2'"`
	Em2_3   uint32  `xorm:"'em2_3'"`
	Em2_4   uint32  `xorm:"'em2_4'"`
	Em2_5   uint32  `xorm:"'em2_5'"`
	Text3_0 string  `xorm:"'text3_0'"`
	Text3_1 string  `xorm:"'text3_1'"`
	Lang3   uint32  `xorm:"'lang3'"`
	Prob3   float32 `xorm:"'prob3'"`
	Em3_0   uint32  `xorm:"'em3_0'"`
	Em3_1   uint32  `xorm:"'em3_1'"`
	Em3_2   uint32  `xorm:"'em3_2'"`
	Em3_3   uint32  `xorm:"'em3_3'"`
	Em3_4   uint32  `xorm:"'em3_4'"`
	Em3_5   uint32  `xorm:"'em3_5'"`
	Text4_0 string  `xorm:"'text4_0'"`
	Text4_1 string  `xorm:"'text4_1'"`
	Lang4   uint32  `xorm:"'lang4'"`
	Prob4   float32 `xorm:"'prob4'"`
	Em4_0   uint32  `xorm:"'em4_0'"`
	Em4_1   uint32  `xorm:"'em4_1'"`
	Em4_2   uint32  `xorm:"'em4_2'"`
	Em4_3   uint32  `xorm:"'em4_3'"`
	Em4_4   uint32  `xorm:"'em4_4'"`
	Em4_5   uint32  `xorm:"'em4_5'"`
	Text5_0 string  `xorm:"'text5_0'"`
	Text5_1 string  `xorm:"'text5_1'"`
	Lang5   uint32  `xorm:"'lang5'"`
	Prob5   float32 `xorm:"'prob5'"`
	Em5_0   uint32  `xorm:"'em5_0'"`
	Em5_1   uint32  `xorm:"'em5_1'"`
	Em5_2   uint32  `xorm:"'em5_2'"`
	Em5_3   uint32  `xorm:"'em5_3'"`
	Em5_4   uint32  `xorm:"'em5_4'"`
	Em5_5   uint32  `xorm:"'em5_5'"`
	Text6_0 string  `xorm:"'text6_0'"`
	Text6_1 string  `xorm:"'text6_1'"`
	Lang6   uint32  `xorm:"'lang6'"`
	Prob6   float32 `xorm:"'prob6'"`
	Em6_0   uint32  `xorm:"'em6_0'"`
	Em6_1   uint32  `xorm:"'em6_1'"`
	Em6_2   uint32  `xorm:"'em6_2'"`
	Em6_3   uint32  `xorm:"'em6_3'"`
	Em6_4   uint32  `xorm:"'em6_4'"`
	Em6_5   uint32  `xorm:"'em6_5'"`
	Text7_0 string  `xorm:"'text7_0'"`
	Text7_1 string  `xorm:"'text7_1'"`
	Lang7   uint32  `xorm:"'lang7'"`
	Prob7   float32 `xorm:"'prob7'"`
	Em7_0   uint32  `xorm:"'em7_0'"`
	Em7_1   uint32  `xorm:"'em7_1'"`
	Em7_2   uint32  `xorm:"'em7_2'"`
	Em7_3   uint32  `xorm:"'em7_3'"`
	Em7_4   uint32  `xorm:"'em7_4'"`
	Em7_5   uint32  `xorm:"'em7_5'"`
}

func (nt NPCText) TableName() string {
	return "npc_text"
}

func (nt NPCText) GetOption(idx int) wdb.NPCTextOption {
	var to wdb.NPCTextOption
	val := reflect.ValueOf(nt)
	if val.IsZero() {
		return to
	}

	to.Text = i18n.GetEnglish(val.FieldByName(fmt.Sprintf("Text%d_0", idx)).String())
	if to.Text == nil {
		to.Text = i18n.GetEnglish(val.FieldByName(fmt.Sprintf("Text%d_1", idx)).String())
	}

	to.Lang = uint32(val.FieldByName(fmt.Sprintf("Lang%d", idx)).Uint())
	to.Prob = float32(val.FieldByName(fmt.Sprintf("Prob%d", idx)).Float())

	for x := 0; x < 6; x += 2 {
		emoteDelay := uint32(val.FieldByName(fmt.Sprintf("Em%d_%d", idx, x)).Uint())

		strName := fmt.Sprintf("Em%d_%d", idx, x+1)
		emote := uint32(val.FieldByName(strName).Uint())

		if emoteDelay != 0 || emote != 0 {
			to.Emote = append(to.Emote, wdb.NPCTextOptionEmote{
				Delay: emoteDelay,
				ID:    emote,
			})
		}
	}

	return to
}

type PlayerCreateInfo struct {
	Race  uint8   `xorm:"'race'"`
	Class uint8   `xorm:"'class'"`
	Map   uint32  `xorm:"'map'"`
	Zone  uint32  `xorm:"'zone'"`
	X     float32 `xorm:"'position_x'"`
	Y     float32 `xorm:"'position_y'"`
	Z     float32 `xorm:"'position_z'"`
	O     float32 `xorm:"'orientation'"`
}

func (pci PlayerCreateInfo) TableName() string {
	return "playercreateinfo"
}

type GameTele struct {
	ID          int64   `xorm:"'id'"`
	PositionX   float32 `xorm:"'position_x'"`
	PositionY   float32 `xorm:"'position_y'"`
	PositionZ   float32 `xorm:"'position_z'"`
	Orientation float32 `xorm:"'orientation'"`
	Map         uint32  `xorm:"'map'"`
	Name        string  `xorm:"'name'"`
}

type AreatriggerTeleport struct {
	ID                uint32  `xorm:"'id'"`
	Name              string  `xorm:"'name'"`
	RequiredLevel     uint32  `xorm:"'required_level'"`
	RequiredItem      uint32  `xorm:"'required_item'"`
	RequiredItem2     uint32  `xorm:"'required_item2'"`
	RequiredQuestDone uint32  `xorm:"'required_quest_done'"`
	TargetMap         uint32  `xorm:"'target_map'"`
	TargetPositionX   float32 `xorm:"'target_position_x'"`
	TargetPositionY   float32 `xorm:"'target_position_y'"`
	TargetPositionZ   float32 `xorm:"'target_position_z'"`
	TargetPositionO   float32 `xorm:"'target_orientation'"`
}

type CreatureTemplate struct {
	Entry                uint32  `xorm:"'Entry'"`
	Name                 string  `xorm:"'Name'"`
	SubName              string  `xorm:"'SubName'"`
	MinLevel             uint32  `xorm:"'MinLevel'"`
	MaxLevel             uint32  `xorm:"'MaxLevel'"`
	ModelId1             uint32  `xorm:"'ModelId1'"`
	ModelId2             uint32  `xorm:"'ModelId2'"`
	ModelId3             uint32  `xorm:"'ModelId3'"`
	ModelId4             uint32  `xorm:"'ModelId4'"`
	Faction              uint32  `xorm:"'Faction'"`
	Scale                float32 `xorm:"'Scale'"`
	Family               int32   `xorm:"'Family'"`
	CreatureType         uint32  `xorm:"'CreatureType'"`
	InhabitType          uint32  `xorm:"'InhabitType'"`
	RegenerateStats      uint32  `xorm:"'RegenerateStats'"`
	RacialLeader         uint32  `xorm:"'RacialLeader'"`
	NpcFlags             uint32  `xorm:"'NpcFlags'"`
	UnitFlags            uint32  `xorm:"'UnitFlags'"`
	DynamicFlags         uint32  `xorm:"'DynamicFlags'"`
	ExtraFlags           uint32  `xorm:"'ExtraFlags'"`
	CreatureTypeFlags    uint32  `xorm:"'CreatureTypeFlags'"`
	SpeedWalk            float32 `xorm:"'SpeedWalk'"`
	SpeedRun             float32 `xorm:"'SpeedRun'"`
	UnitClass            uint32  `xorm:"'UnitClass'"`
	Rank                 uint32  `xorm:"'Rank'"`
	HealthMultiplier     float32 `xorm:"'HealthMultiplier'"`
	PowerMultiplier      float32 `xorm:"'PowerMultiplier'"`
	DamageMultiplier     float32 `xorm:"'DamageMultiplier'"`
	DamageVariance       float32 `xorm:"'DamageVariance'"`
	ArmorMultiplier      float32 `xorm:"'ArmorMultiplier'"`
	ExperienceMultiplier float32 `xorm:"'ExperienceMultiplier'"`
	MinLevelHealth       uint32  `xorm:"'MinLevelHealth'"`
	MaxLevelHealth       uint32  `xorm:"'MaxLevelHealth'"`
	MinLevelMana         uint32  `xorm:"'MinLevelMana'"`
	MaxLevelMana         uint32  `xorm:"'MaxLevelMana'"`
	MinMeleeDmg          float32 `xorm:"'MinMeleeDmg'"`
	MaxMeleeDmg          float32 `xorm:"'MaxMeleeDmg'"`
	MinRangedDmg         float32 `xorm:"'MinRangedDmg'"`
	MaxRangedDmg         float32 `xorm:"'MaxRangedDmg'"`
	Armor                uint32  `xorm:"'Armor'"`
	MeleeAttackPower     uint32  `xorm:"'MeleeAttackPower'"`
	RangedAttackPower    uint32  `xorm:"'RangedAttackPower'"`
	MeleeBaseAttackTime  uint32  `xorm:"'MeleeBaseAttackTime'"`
	RangedBaseAttackTime uint32  `xorm:"'RangedBaseAttackTime'"`
	DamageSchool         int32   `xorm:"'DamageSchool'"`
	MinLootGold          uint32  `xorm:"'MinLootGold'"`
	MaxLootGold          uint32  `xorm:"'MaxLootGold'"`
	LootId               uint32  `xorm:"'LootId'"`
	PickpocketLootId     uint32  `xorm:"'PickpocketLootId'"`
	SkinningLootId       uint32  `xorm:"'SkinningLootId'"`
	KillCredit1          uint32  `xorm:"'KillCredit1'"`
	KillCredit2          uint32  `xorm:"'KillCredit2'"`
	MechanicImmuneMask   uint32  `xorm:"'MechanicImmuneMask'"`
	SchoolImmuneMask     uint32  `xorm:"'SchoolImmuneMask'"`
	ResistanceHoly       int32   `xorm:"'ResistanceHoly'"`
	ResistanceFire       int32   `xorm:"'ResistanceFire'"`
	ResistanceNature     int32   `xorm:"'ResistanceNature'"`
	ResistanceFrost      int32   `xorm:"'ResistanceFrost'"`
	ResistanceShadow     int32   `xorm:"'ResistanceShadow'"`
	ResistanceArcane     int32   `xorm:"'ResistanceArcane'"`
	PetSpellDataId       uint32  `xorm:"'PetSpellDataId'"`
	MovementType         uint32  `xorm:"'MovementType'"`
	TrainerType          int32   `xorm:"'TrainerType'"`
	TrainerSpell         uint32  `xorm:"'TrainerSpell'"`
	TrainerClass         uint32  `xorm:"'TrainerClass'"`
	TrainerRace          uint32  `xorm:"'TrainerRace'"`
	TrainerTemplateId    uint32  `xorm:"'TrainerTemplateId'"`
	VendorTemplateId     uint32  `xorm:"'VendorTemplateId'"`
	GossipMenuId         uint32  `xorm:"'GossipMenuId'"`
	EquipmentTemplateId  uint32  `xorm:"'EquipmentTemplateId'"`
	Civilian             uint32  `xorm:"'Civilian'"`
	AIName               string  `xorm:"'AIName'"`
	ScriptName           string  `xorm:"'ScriptName'"`
}

type GameobjectTemplate struct {
	Entry     uint32 `xorm:"'entry''`
	Type      uint32 `xorm:"'type'"`
	DisplayID uint32 `xorm:"'displayid'"`
	Name      string `xorm:"'name'"`
	// IconName       string `xorm:"'IconName'"`
	// CastBarCaption string `xorm:"'castBarCaption'"`
	Faction     uint32  `xorm:"'faction'"`
	Flags       uint32  `xorm:"'flags'"`
	ExtraFlags  uint32  `xorm:"'ExtraFlags'"`
	Size        float32 `xorm:"'size'"`
	Data0       uint32  `xorm:"'data0'"`
	Data1       uint32  `xorm:"'data1'"`
	Data2       uint32  `xorm:"'data2'"`
	Data3       uint32  `xorm:"'data3'"`
	Data4       uint32  `xorm:"'data4'"`
	Data5       uint32  `xorm:"'data5'"`
	Data6       uint32  `xorm:"'data6'"`
	Data7       uint32  `xorm:"'data7'"`
	Data8       uint32  `xorm:"'data8'"`
	Data9       uint32  `xorm:"'data9'"`
	Data10      uint32  `xorm:"'data10'"`
	Data11      uint32  `xorm:"'data11'"`
	Data12      uint32  `xorm:"'data12'"`
	Data13      uint32  `xorm:"'data13'"`
	Data14      uint32  `xorm:"'data14'"`
	Data15      uint32  `xorm:"'data15'"`
	Data16      uint32  `xorm:"'data16'"`
	Data17      uint32  `xorm:"'data17'"`
	Data18      uint32  `xorm:"'data18'"`
	Data19      uint32  `xorm:"'data19'"`
	Data20      uint32  `xorm:"'data20'"`
	Data21      uint32  `xorm:"'data21'"`
	Data22      uint32  `xorm:"'data22'"`
	Data23      uint32  `xorm:"'data23'"`
	CustomData1 uint32  `xorm:"'CustomData1'"`
	Mingold     int64   `xorm:"'mingold'"`
	Maxgold     int64   `xorm:"'maxgold'"`
	ScriptName  string  `xorm:"'ScriptName'"`
}

type ItemTemplate struct {
	Entry                     int    `xorm:"'entry'"`
	Class                     uint32 `xorm:"'class'"`
	Subclass                  uint32 `xorm:"'subclass'"`
	Name                      string `xorm:"'name'"`
	DisplayID                 uint32 `xorm:"'displayid'"`
	Quality                   uint8  `xorm:"'Quality'"`
	Flags                     uint32 `xorm:"'Flags'"`
	BuyCount                  uint8  `xorm:"'BuyCount'"`
	BuyPrice                  uint32 `xorm:"'BuyPrice'"`
	SellPrice                 uint32 `xorm:"'SellPrice'"`
	InventoryType             uint8  `xorm:"'InventoryType'"`
	AllowableClass            int32  `xorm:"'AllowableClass'"`
	AllowableRace             int32  `xorm:"'AllowableRace'"`
	ItemLevel                 uint32 `xorm:"'ItemLevel'"`
	RequiredLevel             uint8  `xorm:"'RequiredLevel'"`
	RequiredSkill             uint32 `xorm:"'RequiredSkill'"`
	RequiredSkillRank         uint32 `xorm:"'RequiredSkillRank'"`
	Requiredspell             uint32 `xorm:"'requiredspell'"`
	Requiredhonorrank         uint32 `xorm:"'requiredhonorrank'"`
	RequiredCityRank          uint32 `xorm:"'RequiredCityRank'"`
	RequiredReputationFaction uint32 `xorm:"'RequiredReputationFaction'"`
	RequiredReputationRank    uint32 `xorm:"'RequiredReputationRank'"`
	Maxcount                  uint32 `xorm:"'maxcount'"`
	Stackable                 uint32 `xorm:"'stackable'"`
	ContainerSlots            uint8  `xorm:"'ContainerSlots'"`
	// StatsCount                uint8   `xorm:"'StatsCount'"`
	Stat_type1   uint8 `xorm:"'stat_type1'"`
	Stat_value1  int32 `xorm:"'stat_value1'"`
	Stat_type2   uint8 `xorm:"'stat_type2'"`
	Stat_value2  int32 `xorm:"'stat_value2'"`
	Stat_type3   uint8 `xorm:"'stat_type3'"`
	Stat_value3  int32 `xorm:"'stat_value3'"`
	Stat_type4   uint8 `xorm:"'stat_type4'"`
	Stat_value4  int32 `xorm:"'stat_value4'"`
	Stat_type5   uint8 `xorm:"'stat_type5'"`
	Stat_value5  int32 `xorm:"'stat_value5'"`
	Stat_type6   uint8 `xorm:"'stat_type6'"`
	Stat_value6  int32 `xorm:"'stat_value6'"`
	Stat_type7   uint8 `xorm:"'stat_type7'"`
	Stat_value7  int32 `xorm:"'stat_value7'"`
	Stat_type8   uint8 `xorm:"'stat_type8'"`
	Stat_value8  int32 `xorm:"'stat_value8'"`
	Stat_type9   uint8 `xorm:"'stat_type9'"`
	Stat_value9  int32 `xorm:"'stat_value9'"`
	Stat_type10  uint8 `xorm:"'stat_type10'"`
	Stat_value10 int32 `xorm:"'stat_value10'"`
	// ScalingStatDistribution int32   `xorm:"'ScalingStatDistribution'"`
	// ScalingStatValue        int32   `xorm:"'ScalingStatValue'"`
	DMG_min1                float32 `xorm:"'dmg_min1'"`
	DMG_max1                float32 `xorm:"'dmg_max1'"`
	DMG_type1               uint8   `xorm:"'dmg_type1'"`
	DMG_min2                float32 `xorm:"'dmg_min2'"`
	DMG_max2                float32 `xorm:"'dmg_max2'"`
	DMG_type2               uint8   `xorm:"'dmg_type2'"`
	DMG_min3                float32 `xorm:"'dmg_min3'"`
	DMG_max3                float32 `xorm:"'dmg_max3'"`
	DMG_type3               uint8   `xorm:"'dmg_type3'"`
	DMG_min4                float32 `xorm:"'dmg_min4'"`
	DMG_max4                float32 `xorm:"'dmg_max4'"`
	DMG_type4               uint8   `xorm:"'dmg_type4'"`
	DMG_min5                float32 `xorm:"'dmg_min5'"`
	DMG_max5                float32 `xorm:"'dmg_max5'"`
	DMG_type5               uint8   `xorm:"'dmg_type5'"`
	Armor                   uint32  `xorm:"'armor'"`
	Holy_res                uint32  `xorm:"'holy_res'"`
	Fire_res                uint32  `xorm:"'fire_res'"`
	Nature_res              uint32  `xorm:"'nature_res'"`
	Frost_res               uint32  `xorm:"'frost_res'"`
	Shadow_res              uint32  `xorm:"'shadow_res'"`
	Arcane_res              uint32  `xorm:"'arcane_res'"`
	Delay                   uint32  `xorm:"'delay'"`
	Ammo_type               uint32  `xorm:"'ammo_type'"`
	RangedModRange          float32 `xorm:"'RangedModRange'"`
	Spellid_1               uint32  `xorm:"'spellid_1'"`
	Spelltrigger_1          uint32  `xorm:"'spelltrigger_1'"`
	Spellcharges_1          int     `xorm:"'spellcharges_1'"`
	SpellppmRate_1          float32 `xorm:"'spellppmRate_1'"`
	Spellcooldown_1         int32   `xorm:"'spellcooldown_1'"`
	Spellcategory_1         uint32  `xorm:"'spellcategory_1'"`
	Spellcategorycooldown_1 int32   `xorm:"'spellcategorycooldown_1'"`
	Spellid_2               uint32  `xorm:"'spellid_2'"`
	Spelltrigger_2          uint32  `xorm:"'spelltrigger_2'"`
	Spellcharges_2          int32   `xorm:"'spellcharges_2'"`
	SpellppmRate_2          float32 `xorm:"'spellppmRate_2'"`
	Spellcooldown_2         int32   `xorm:"'spellcooldown_2'"`
	Spellcategory_2         uint32  `xorm:"'spellcategory_2'"`
	Spellcategorycooldown_2 int32   `xorm:"'spellcategorycooldown_2'"`
	Spellid_3               uint32  `xorm:"'spellid_3'"`
	Spelltrigger_3          uint32  `xorm:"'spelltrigger_3'"`
	Spellcharges_3          int32   `xorm:"'spellcharges_3'"`
	SpellppmRate_3          float32 `xorm:"'spellppmRate_3'"`
	Spellcooldown_3         int32   `xorm:"'spellcooldown_3'"`
	Spellcategory_3         uint32  `xorm:"'spellcategory_3'"`
	Spellcategorycooldown_3 int32   `xorm:"'spellcategorycooldown_3'"`
	Spellid_4               uint32  `xorm:"'spellid_4'"`
	Spelltrigger_4          uint32  `xorm:"'spelltrigger_4'"`
	Spellcharges_4          int32   `xorm:"'spellcharges_4'"`
	SpellppmRate_4          float32 `xorm:"'spellppmRate_4'"`
	Spellcooldown_4         int32   `xorm:"'spellcooldown_4'"`
	Spellcategory_4         uint32  `xorm:"'spellcategory_4'"`
	Spellcategorycooldown_4 int32   `xorm:"'spellcategorycooldown_4'"`
	Spellid_5               uint32  `xorm:"'spellid_5'"`
	Spelltrigger_5          uint32  `xorm:"'spelltrigger_5'"`
	Spellcharges_5          int32   `xorm:"'spellcharges_5'"`
	SpellppmRate_5          float32 `xorm:"'spellppmRate_5'"`
	Spellcooldown_5         int32   `xorm:"'spellcooldown_5'"`
	Spellcategory_5         uint32  `xorm:"'spellcategory_5'"`
	Spellcategorycooldown_5 int32   `xorm:"'spellcategorycooldown_5'"`
	Bonding                 uint8   `xorm:"'bonding'"`
	Description             string  `xorm:"'description'"`
	PageText                uint32  `xorm:"'PageText'"`
	LanguageID              uint32  `xorm:"'LanguageID'"`
	PageMaterial            uint32  `xorm:"'PageMaterial'"`
	Startquest              uint32  `xorm:"'startquest'"`
	Lockid                  uint32  `xorm:"'lockid'"`
	Material                int32   `xorm:"'Material'"`
	Sheath                  uint32  `xorm:"'sheath'"`
	RandomProperty          uint32  `xorm:"'RandomProperty'"`
	// RandomSuffix            uint32  `xorm:"'RandomSuffix'"`
	Block         uint32 `xorm:"'block'"`
	Itemset       uint32 `xorm:"'itemset'"`
	MaxDurability uint32 `xorm:"'MaxDurability'"`
	Area          uint32 `xorm:"'area'"`
	Map           int32  `xorm:"'Map'"`
	BagFamily     int32  `xorm:"'BagFamily'"`
	// TotemCategory           int32   `xorm:"'TotemCategory'"`
	// SocketColor_1           int32   `xorm:"'socketColor_1'"`
	// SocketContent_1         int32   `xorm:"'socketContent_1'"`
	// SocketColor_2           int32   `xorm:"'socketColor_2'"`
	// SocketContent_2         int32   `xorm:"'socketContent_2'"`
	// SocketColor_3           int32   `xorm:"'socketColor_3'"`
	// SocketContent_3         int32   `xorm:"'socketContent_3'"`
	// SocketBonus             int32   `xorm:"'socketBonus'"`
	// GemProperties           int32   `xorm:"'GemProperties'"`
	// RequiredDisenchantSkill int32   `xorm:"'RequiredDisenchantSkill'"`
	// ArmorDamageModifier float32 `xorm:"'ArmorDamageModifier'"`
	// ItemLimitCategory int32  `xorm:"'ItemLimitCategory'"`
	ScriptName   string `xorm:"'ScriptName'"`
	DisenchantID uint32 `xorm:"'DisenchantID'"`
	FoodType     uint8  `xorm:"'FoodType'"`
	MinMoneyLoot uint32 `xorm:"'minMoneyLoot'"`
	MaxMoneyLoot uint32 `xorm:"'maxMoneyLoot'"`
	Duration     int32  `xorm:"'Duration'"`
	ExtraFlags   uint8  `xorm:"'ExtraFlags'"`
}

func (it ItemTemplate) GetItemStat(idx int) wdb.ItemStat {
	st := fmt.Sprintf("Stat_type%d", idx)
	sv := fmt.Sprintf("Stat_value%d", idx)
	val := reflect.ValueOf(it)
	_, ok := val.Type().FieldByName(st)
	if !ok {
		panic("no field for " + st)
	}

	return wdb.ItemStat{
		uint8(val.FieldByName(st).Uint()),
		int32(val.FieldByName(sv).Int()),
	}
}

func (it ItemTemplate) GetItemDamage(idx int) wdb.ItemDamage {
	dt := fmt.Sprintf("DMG_type%d", idx)
	dmn := fmt.Sprintf("DMG_min%d", idx)
	dmx := fmt.Sprintf("DMG_max%d", idx)

	val := reflect.ValueOf(it)
	_, ok := val.Type().FieldByName(dt)
	if !ok {
		panic("no field for " + dt)
	}

	return wdb.ItemDamage{
		uint8(val.FieldByName(dt).Uint()),
		float32(val.FieldByName(dmn).Float()),
		float32(val.FieldByName(dmx).Float()),
	}
}

func (it ItemTemplate) GetItemSpell(idx int) wdb.ItemSpell {
	st := fmt.Sprintf("Spellid_%d", idx)
	strigger := fmt.Sprintf("Spelltrigger_%d", idx)
	scharges := fmt.Sprintf("Spellcharges_%d", idx)
	ppm := fmt.Sprintf("SpellppmRate_%d", idx)
	cool := fmt.Sprintf("Spellcooldown_%d", idx)
	scategory := fmt.Sprintf("Spellcategory_%d", idx)
	scategoryCooldown := fmt.Sprintf("Spellcategorycooldown_%d", idx)

	val := reflect.ValueOf(it)
	_, ok := val.Type().FieldByName(st)
	if !ok {
		panic("no field for " + st)
	}

	return wdb.ItemSpell{
		uint32(val.FieldByName(st).Uint()),
		uint32(val.FieldByName(strigger).Uint()),
		int32(val.FieldByName(scharges).Int()),
		float32(val.FieldByName(ppm).Float()),
		int64(val.FieldByName(cool).Int()),
		uint32(val.FieldByName(scategory).Uint()),
		int64(val.FieldByName(scategoryCooldown).Int()),
	}
}

func printTimestamp(out io.Writer) {
	fmt.Fprintf(out, "// DO NOT EDIT: extracted from CMaNGOS database on %s\n", datapack.Timestamp())
}

func openTextWriter(out io.Writer) *text.Encoder {
	j := text.NewEncoder(out)
	return j
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(os.Args[0], "<xorm url>")
		return
	}

	c, err := xorm.NewEngine("mysql", os.Args[1])
	if err != nil {
		panic(err)
	}

	var gt []GameTele
	err = c.Find(&gt)
	if err != nil {
		panic(err)
	}

	plc := openFile("DB/PortLocation.txt")

	printTimestamp(plc)
	pwr := openTextWriter(plc)

	for _, pl := range gt {
		if err := pwr.Encode(wdb.PortLocation{
			ID:  pl.Name,
			X:   pl.PositionX,
			Y:   pl.PositionY,
			Z:   pl.PositionZ,
			O:   pl.Orientation,
			Map: pl.Map,
		}); err != nil {
			panic(err)
		}
	}

	plc.Close()

	var itt []ItemTemplate
	err = c.Find(&itt)
	if err != nil {
		panic(err)
	}

	fl := openFile("DB/NPCText.txt")
	wr := text.NewEncoder(fl)

	var npcText []NPCText

	err = c.Find(&npcText)
	if err != nil {
		panic(err)
	}

	for _, text := range npcText {
		var nt wdb.NPCText

		nt.ID = fmt.Sprintf("nt:%d", text.ID)
		for x := 0; x < 8; x++ {
			opt := text.GetOption(x)

			if reflect.ValueOf(opt).IsZero() == false {
				nt.Opts = append(nt.Opts, opt)
			}
		}

		if err := wr.Encode(nt); err != nil {
			panic(err)
		}
	}

	fl.Close()

	fl = openFile("DB/ItemTemplate.txt")

	printTimestamp(fl)

	wr = openTextWriter(fl)

	for _, t := range itt {
		var stats []wdb.ItemStat

		for i := 0; i < 10; i++ {
			st := t.GetItemStat(i + 1)
			if st.Type != 0 {
				stats = append(stats, st)
			}
		}

		var dmg []wdb.ItemDamage
		for i := 0; i < 5; i++ {
			d := t.GetItemDamage(i + 1)
			if d.Type != 0 {
				dmg = append(dmg, d)
			}
		}

		var spell []wdb.ItemSpell
		for i := 0; i < 5; i++ {
			sp := t.GetItemSpell(i + 1)
			if sp.ID != 0 {
				spell = append(spell, sp)
			}
		}

		flags, err := update.DecodeItemFlagInteger(5875, uint64(t.Flags))
		if err != nil {
			panic(err)
		}

		witem := wdb.ItemTemplate{
			ID:                        fmt.Sprintf("it:%d", t.Entry),
			Class:                     t.Class,
			Subclass:                  t.Subclass,
			Name:                      i18n.GetEnglish(t.Name),
			DisplayID:                 t.DisplayID,
			Quality:                   t.Quality,
			Flags:                     flags.String(), //todo: convert flags to a readable form
			BuyCount:                  t.BuyCount,
			BuyPrice:                  t.BuyPrice,
			SellPrice:                 t.SellPrice,
			InventoryType:             t.InventoryType,
			AllowableClass:            t.AllowableClass,
			AllowableRace:             t.AllowableRace,
			ItemLevel:                 t.ItemLevel,
			RequiredLevel:             t.RequiredLevel,
			RequiredSkill:             t.RequiredSkill,
			RequiredSkillRank:         t.RequiredSkillRank,
			RequiredSpell:             t.Requiredspell,
			RequiredHonorRank:         t.Requiredhonorrank,
			RequiredCityRank:          t.RequiredCityRank,
			RequiredReputationFaction: t.RequiredReputationFaction,
			MaxCount:                  t.Maxcount,
			Stackable:                 t.Stackable,
			ContainerSlots:            t.ContainerSlots,
			Stats:                     stats,
			Damage:                    dmg,
			Armor:                     t.Armor,
			HolyRes:                   t.Holy_res,
			FireRes:                   t.Fire_res,
			NatureRes:                 t.Nature_res,
			FrostRes:                  t.Frost_res,
			ShadowRes:                 t.Shadow_res,
			ArcaneRes:                 t.Arcane_res,
			Delay:                     t.Delay,
			AmmoType:                  t.Ammo_type,
			RangedModRange:            t.RangedModRange,
			Spells:                    spell,
			Bonding:                   t.Bonding,
			Description:               i18n.GetEnglish(t.Description),
			PageText:                  t.PageText,
			LanguageID:                t.LanguageID,
			PageMaterial:              t.PageMaterial,
			StartQuest:                t.Startquest,
			LockID:                    t.Lockid,
			Material:                  t.Material,
			Sheath:                    t.Sheath,
			RandomProperty:            t.RandomProperty,
			Block:                     t.Block,
			Itemset:                   t.Itemset,
			MaxDurability:             t.MaxDurability,
			Area:                      t.Area,
			Map:                       t.Map,
			BagFamily:                 t.BagFamily,
			ScriptName:                t.ScriptName,
			DisenchantID:              t.DisenchantID,
			FoodType:                  t.FoodType,
			MinMoneyLoot:              econ.Money(t.MinMoneyLoot),
			MaxMoneyLoot:              econ.Money(t.MaxMoneyLoot),
			Duration:                  t.Duration,
			ExtraFlags:                t.ExtraFlags,
		}
		if err := wr.Encode(witem); err != nil {
			panic(err)
		}
	}

	fl.Close()

	itt = nil

	var gtt []GameobjectTemplate
	err = c.Find(&gtt)
	if err != nil {
		panic(err)
	}

	gfl := openFile("DB/GameObjectTemplate.txt")
	printTimestamp(gfl)
	wr = openTextWriter(gfl)

	for _, v := range gtt {
		data := make([]uint32, 24)
		d := reflect.ValueOf(v)
		for x := 0; x < 24; x++ {
			data[x] = uint32(d.FieldByName(fmt.Sprintf("Data%d", x)).Uint())
		}

		wr.Encode(wdb.GameObjectTemplate{
			ID:        fmt.Sprintf("go:%d", v.Entry),
			Type:      v.Type,
			DisplayID: v.DisplayID,
			Name:      i18n.GetEnglish(v.Name),
			// IconName:       v.IconName,
			// CastBarCaption: v.CastBarCaption,
			Faction:       v.Faction,
			Flags:         update.GameObjectFlags(v.Flags).String(),
			HasCustomAnim: v.ExtraFlags == 1,
			Size:          v.Size,
			Data:          data,
			MinGold:       econ.Money(v.Mingold),
			MaxGold:       econ.Money(v.Maxgold),
		})
	}

	gfl.Close()

	gtt = nil

	var ctt []CreatureTemplate
	err = c.Find(&ctt)
	if err != nil {
		panic(err)
	}

	cfl := openFile("DB/CreatureTemplate.txt")
	wr = openTextWriter(cfl)
	for _, cr := range ctt {
		ct := wdb.CreatureTemplate{
			ID:                  fmt.Sprintf("cr:%d", cr.Entry),
			Entry:               cr.Entry,
			Name:                i18n.GetEnglish(cr.Name),
			SubName:             i18n.GetEnglish(cr.SubName),
			MinLevel:            cr.MinLevel,
			MaxLevel:            cr.MaxLevel,
			Faction:             cr.Faction,
			Scale:               cr.Scale,
			CreatureType:        cr.CreatureType,
			InhabitType:         cr.InhabitType,
			RegenerateStats:     cr.RegenerateStats,
			RacialLeader:        cr.RacialLeader == 1,
			Gossip:              cr.NpcFlags&1 != 0,
			QuestGiver:          cr.NpcFlags&2 != 0,
			Vendor:              cr.NpcFlags&4 != 0,
			FlightMaster:        cr.NpcFlags&8 != 0,
			Trainer:             cr.NpcFlags&16 != 0,
			SpiritHealer:        cr.NpcFlags&32 != 0,
			SpiritGuide:         cr.NpcFlags&64 != 0,
			Innkeeper:           cr.NpcFlags&128 != 0,
			Banker:              cr.NpcFlags&256 != 0,
			Petitioner:          cr.NpcFlags&512 != 0,
			TabardDesigner:      cr.NpcFlags&1024 != 0,
			BattleMaster:        cr.NpcFlags&2048 != 0,
			Auctioneer:          cr.NpcFlags&4096 != 0,
			StableMaster:        cr.NpcFlags&8192 != 0,
			Repairer:            cr.NpcFlags&16384 != 0,
			OutdoorPVP:          cr.NpcFlags&536870912 != 0,
			ServerControlled:    cr.UnitFlags&0x1 != 0,
			NonAttackable:       cr.UnitFlags&0x2 != 0,
			RemoveClientControl: cr.UnitFlags&0x4 != 0,
			PlayerControlled:    cr.UnitFlags&0x8 != 0,
			Rename:              cr.UnitFlags&0x10 != 0,
			PetAbandon:          cr.UnitFlags&0x20 != 0,
			OOCNotAttackable:    cr.UnitFlags&0x100 != 0,
			Passive:             cr.UnitFlags&0x200 != 0,
			PVP:                 cr.UnitFlags&0x1000 != 0,
			IsSilenced:          cr.UnitFlags&0x2000 != 0,
			IsPersuaded:         cr.UnitFlags&0x4000 != 0,
			Swimming:            cr.UnitFlags&0x8000 != 0,
			RemoveAttackIcon:    cr.UnitFlags&0x10000 != 0,
			IsPacified:          cr.UnitFlags&0x20000 != 0,
			IsStunned:           cr.UnitFlags&0x40000 != 0,
			InCombat:            cr.UnitFlags&0x80000 != 0,
			InTaxiFlight:        cr.UnitFlags&0x100000 != 0,
			Disarmed:            cr.UnitFlags&0x200000 != 0,
			Confused:            cr.UnitFlags&0x400000 != 0,
			Fleeing:             cr.UnitFlags&0x800000 != 0,
			Possessed:           cr.UnitFlags&0x1000000 != 0,
			NotSelectable:       cr.UnitFlags&0x2000000 != 0,
			Skinnable:           cr.UnitFlags&0x4000000 != 0,
			AurasVisible:        cr.UnitFlags&0x8000000 != 0,
			Sheathe:             cr.UnitFlags&0x40000000 != 0,
			NoKillReward:        cr.UnitFlags&0x80000000 != 0,

			Lootable:              cr.DynamicFlags&1 != 0,
			TrackUnit:             cr.DynamicFlags&2 != 0,
			Tapped:                cr.DynamicFlags&4 != 0,
			TappedByPlayer:        cr.DynamicFlags&8 != 0,
			SpecialInfo:           cr.DynamicFlags&16 != 0,
			VisuallyDead:          cr.DynamicFlags&32 != 0,
			TappedByAllThreatList: cr.DynamicFlags&128 != 0,

			InstanceBind:             cr.ExtraFlags&0x1 != 0,      // creature kill bind instance with killer and killer’s group
			NoAggroOnSight:           cr.ExtraFlags&0x2 != 0,      // no aggro (ignore faction/reputation hostility)
			NoParry:                  cr.ExtraFlags&0x4 != 0,      // creature can’t parry
			NoParryHasten:            cr.ExtraFlags&0x8 != 0,      // creature can’t counter-attack at parry
			NoBlock:                  cr.ExtraFlags&0x10 != 0,     //	creature can’t block
			NoCrush:                  cr.ExtraFlags&0x20 != 0,     // creature can’t do crush attacks
			NoXPAtKill:               cr.ExtraFlags&0x40 != 0,     // creature kill not provide XP
			Invisible:                cr.ExtraFlags&0x80 != 0,     // creature is always invisible for player (mostly trigger creatures)
			NotTauntable:             cr.ExtraFlags&0x100 != 0,    // creature is immune to taunt auras and effect attack me
			AggroZone:                cr.ExtraFlags&0x200 != 0,    // creature sets itself in combat with zone on aggro
			Guard:                    cr.ExtraFlags&0x400 != 0,    // creature is a guard
			NoCallAssist:             cr.ExtraFlags&0x800 != 0,    // creature shouldn’t call for assistance on aggro
			Active:                   cr.ExtraFlags&0x1000 != 0,   //creature is active object. Grid of this creature will be loaded and creature set as active
			ForceEnableMMap:          cr.ExtraFlags&0x2000 != 0,   // creature is forced to use MMaps
			ForceDisableMMap:         cr.ExtraFlags&0x4000 != 0,   // creature is forced to NOT use MMaps
			WalkInWater:              cr.ExtraFlags&0x8000 != 0,   // creature is forced to walk in water even it can swim
			Civilian:                 cr.ExtraFlags&0x10000 != 0,  // CreatureInfo→civilian substitute (for expansions as Civilian Colum was removed)
			NoMelee:                  cr.ExtraFlags&0x20000 != 0,  // creature can’t melee
			FarView:                  cr.ExtraFlags&0x40000 != 0,  // creature with far view
			ForceAttackingCapability: cr.ExtraFlags&0x80000 != 0,  // SetForceAttackingCapability(true); for nonattackable, nontargetable creatures that should be able to attack nontheless
			IgnoreUsedPosition:       cr.ExtraFlags&0x100000 != 0, // ignore creature when checking used positions around target
			CountSpawns:              cr.ExtraFlags&0x200000 != 0, // count creature spawns in Map*
			HasteSpellImmunity:       cr.ExtraFlags&0x400000 != 0, // immunity to COT or Mind Numbing Poison – very common in instances

			Tameable:                cr.CreatureTypeFlags&1 != 0, // Makes the mob tameable (must also be a beast and have family set)
			VisibleToGhosts:         cr.CreatureTypeFlags&2 != 0, // Sets Creatures that can ALSO be seen when player is a ghost. Used in CanInteract function by client, can’t be attacked
			BossLevel:               cr.CreatureTypeFlags&4 != 0,
			DontPlayWoundParryAnim:  cr.CreatureTypeFlags&8 != 0,
			HideFactionTooltip:      cr.CreatureTypeFlags&16 != 0, // Controls something in client tooltip related to creature faction
			SpellAttackable:         cr.CreatureTypeFlags&64 != 0,
			DeadInteract:            cr.CreatureTypeFlags&128 != 0,
			HerbLoot:                cr.CreatureTypeFlags&256 != 0, // Uses Skinning Loot Field
			MiningLoot:              cr.CreatureTypeFlags&512 != 0, // Makes Mob Corpse Mineable – Uses Skinning Loot Field
			DontLogDeath:            cr.CreatureTypeFlags&1024 != 0,
			MountedCombat:           cr.CreatureTypeFlags&2048 != 0,
			CanAssist:               cr.CreatureTypeFlags&4096 != 0, //	Can aid any player or group in combat. Typically seen for escorting NPC’s
			PetHasActionBar:         cr.CreatureTypeFlags&8192 != 0, // 	checked from calls in Lua_PetHasActionBar
			MaskUID:                 cr.CreatureTypeFlags&16384 != 0,
			EngineerLoot:            cr.CreatureTypeFlags&32768 != 0, //	Makes Mob Corpse Engineer Lootable – Uses Skinning Loot Field
			ExoticPet:               cr.CreatureTypeFlags&65536 != 0, // Tamable as an exotic pet. Normal tamable flag must also be set.
			UseDefaultCollisionBox:  cr.CreatureTypeFlags&131072 != 0,
			IsSiegeWeapon:           cr.CreatureTypeFlags&262144 != 0,
			ProjectileCollision:     cr.CreatureTypeFlags&524288 != 0,
			HideNameplate:           cr.CreatureTypeFlags&1048576 != 0,
			DontPlayMountedAnim:     cr.CreatureTypeFlags&2097152 != 0,
			IsLinkAll:               cr.CreatureTypeFlags&4194304 != 0,
			InteractOnlyWithCreator: cr.CreatureTypeFlags&8388608 != 0,
			ForceGossip:             cr.CreatureTypeFlags&134217728 != 0,

			SpeedWalk:            cr.SpeedWalk,
			SpeedRun:             cr.SpeedRun,
			UnitClass:            cr.UnitClass,
			Rank:                 cr.Rank,
			HealthMultiplier:     cr.HealthMultiplier,
			PowerMultiplier:      cr.PowerMultiplier,
			DamageMultiplier:     cr.DamageMultiplier,
			DamageVariance:       cr.DamageVariance,
			ArmorMultiplier:      cr.ArmorMultiplier,
			ExperienceMultiplier: cr.ExperienceMultiplier,
			MinLevelHealth:       cr.MinLevelHealth,
			MaxLevelHealth:       cr.MaxLevelHealth,
			MinLevelMana:         cr.MinLevelMana,
			MaxLevelMana:         cr.MaxLevelMana,
			MinMeleeDmg:          cr.MinMeleeDmg,
			MaxMeleeDmg:          cr.MaxMeleeDmg,
			MinRangedDmg:         cr.MinRangedDmg,
			MaxRangedDmg:         cr.MaxRangedDmg,
			Armor:                cr.Armor,
			MeleeAttackPower:     cr.MeleeAttackPower,
			RangedAttackPower:    cr.RangedAttackPower,
			MeleeBaseAttackTime:  cr.MeleeBaseAttackTime,
			RangedBaseAttackTime: cr.RangedBaseAttackTime,
			DamageSchool:         cr.DamageSchool,
			MinLootGold:          econ.Money(cr.MinLootGold),
			MaxLootGold:          econ.Money(cr.MaxLootGold),
			LootId:               cr.LootId,
			PickpocketLootId:     cr.PickpocketLootId,
			SkinningLootId:       cr.SkinningLootId,
			KillCredit1:          cr.KillCredit1,
			KillCredit2:          cr.KillCredit2,
			MechanicImmuneMask:   cr.MechanicImmuneMask,
			SchoolImmuneMask:     cr.SchoolImmuneMask,
			ResistanceHoly:       cr.ResistanceHoly,
			ResistanceFire:       cr.ResistanceFire,
			ResistanceNature:     cr.ResistanceNature,
			ResistanceFrost:      cr.ResistanceFrost,
			ResistanceShadow:     cr.ResistanceShadow,
			ResistanceArcane:     cr.ResistanceArcane,
			PetSpellDataId:       cr.PetSpellDataId,
			MovementType:         cr.MovementType,
			TrainerType:          cr.TrainerType,
			TrainerSpell:         cr.TrainerSpell,
			TrainerClass:         cr.TrainerClass,
			TrainerRace:          cr.TrainerRace,
			TrainerTemplateId:    cr.TrainerTemplateId,
			VendorTemplateId:     cr.VendorTemplateId,
			GossipMenuId:         fmt.Sprintf("%d", cr.GossipMenuId),
			EquipmentTemplateId:  cr.EquipmentTemplateId,
			DishonourableKill:    cr.Civilian == 1,
			AIName:               cr.AIName,
			ScriptName:           cr.ScriptName,
		}

		if cr.ModelId1 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId1)
		}

		if cr.ModelId2 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId2)
		}

		if cr.ModelId3 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId3)
		}

		if cr.ModelId4 != 0 {
			ct.DisplayIDs = append(ct.DisplayIDs, cr.ModelId4)
		}

		if cr.Family != 0 {
			switch cr.Family {
			case 1:
				ct.Family = "Wolf"
			case 2:
				ct.Family = "Cat"
			case 3:
				ct.Family = "Spider"
			case 4:
				ct.Family = "Bear"
			case 5:
				ct.Family = "Boar"
			case 6:
				ct.Family = "Crocolisk"
			case 7:
				ct.Family = "Carrion Bird"
			case 8:
				ct.Family = "Crab"
			case 9:
				ct.Family = "Gorilla"
			// case 10:
			// 	ct.Family = "Horse (Custom)"
			case 11:
				ct.Family = "Raptor"
			case 12:
				ct.Family = "Tallstrider"
			case 15:
				ct.Family = "Felhunter"
			case 16:
				ct.Family = "Voidwalker"
			case 17:
				ct.Family = "Succubus"
			case 19:
				ct.Family = "Doomguard"
			case 20:
				ct.Family = "Scorpid"
			case 21:
				ct.Family = "Turtle"
			case 23:
				ct.Family = "Imp"
			case 24:
				ct.Family = "Bat"
			case 25:
				ct.Family = "Hyena"
			case 26:
				ct.Family = "Owl"
			case 27:
				ct.Family = "Wind Serpent"
			case 28:
				ct.Family = "Remote Control"
			default:
				yo.Spew(cr)
				panic(fmt.Errorf("unknown id: %d", cr.Family))
			}
		}

		if err := wr.Encode(ct); err != nil {
			panic(err)
		}
	}

	cfl.Close()

	var att []AreatriggerTeleport

	err = c.Find(&att)
	if err != nil {
		panic(err)
	}

	fl = openFile("Scripts/AreaTriggers.lua")
	for _, v := range att {
		fmt.Fprintf(fl, "-- %s\n", v.Name)
		fmt.Fprintf(fl, "OnAreaTrigger(%d, function(plyr)\n", v.ID)
		if v.RequiredItem != 0 {
			addItemReq(fl, v.RequiredItem)
		}

		if v.RequiredItem2 != 0 {
			addItemReq(fl, v.RequiredItem2)
		}

		if v.RequiredLevel != 0 {
			fmt.Fprintf(fl, "  if plyr:GetLevel() < %d then\n", v.RequiredLevel)
			fmt.Fprintf(fl, "    plyr:SendRequiredLevelZoneError(%d)\n", v.RequiredLevel)
			fmt.Fprintf(fl, "    return\n")
			fmt.Fprintf(fl, "  end\n\n")
		}

		if v.RequiredQuestDone != 0 {
			fmt.Fprintf(fl, "  if not plyr:QuestDone(%d) then\n", v.RequiredQuestDone)
			fmt.Fprintf(fl, "    plyr:SendRequiredQuestZoneError(%d)\n", v.RequiredQuestDone)
			fmt.Fprintf(fl, "    return\n")
			fmt.Fprintf(fl, "  end\n\n")
		}

		fmt.Fprintf(fl, "  plyr:Teleport(%d, %f, %f, %f, %f)\n", v.TargetMap, v.TargetPositionX, v.TargetPositionY, v.TargetPositionZ, v.TargetPositionO)
		fmt.Fprintf(fl, "end)\n\n")
	}

	fl.Close()

	fl = openFile("DB/PlayerCreateInfo.txt")
	wr = text.NewEncoder(fl)

	var pCreateInfo []PlayerCreateInfo

	err = c.Find(&pCreateInfo)
	if err != nil {
		panic(err)
	}

	for _, pci := range pCreateInfo {
		if err := wr.Encode(pci); err != nil {
			panic(err)
		}
	}

	fl.Close()

}

func addItemReq(fl *os.File, item uint32) {
	fmt.Fprintf(fl, "  if not plyr:HasItem(\"it:%d\") then\n", item)
	fmt.Fprintf(fl, "    plyr:SendRequiredItemZoneError(\"it:%d\")\n", item)
	fmt.Fprintf(fl, "    return\n")
	fmt.Fprintf(fl, "  end\n\n")
}

func openFile(out string) *os.File {
	fmt.Println("Extracting to", out, "...")
	fl, _ := os.OpenFile(out, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0700)
	return fl
}
