package sc_play

import (
	"Packetor/packetor/decode"
	error2 "Packetor/packetor/error"
	"errors"
	"fmt"
)

type UpdateRecipesIngredient struct {
	Items []decode.Slot
}

type UpdateRecipesCraftingShapeless struct {
	Type        string
	Group       string
	Category    int32
	Ingredients []UpdateRecipesIngredient
	Result      decode.Slot
}

type UpdateRecipesCraftingShaped struct {
	Type             string
	Width            int32
	Height           int32
	Group            string
	Category         int32
	Ingredients      []UpdateRecipesIngredient
	Result           decode.Slot
	ShowNotification bool
}

type UpdateRecipesCraftingSpecial struct {
	Type     string
	Category int32
}

type UpdateRecipesSmelting struct {
	Type        string
	Group       string
	Category    int32
	Ingredient  UpdateRecipesIngredient
	Result      decode.Slot
	Experience  float32
	CookingTime int32
}

type UpdateRecipesCutting struct {
	Type       string
	Group      string
	Ingredient UpdateRecipesIngredient
	Result     decode.Slot
}

type UpdateRecipesSmithingTransform struct {
	Type     string
	Template UpdateRecipesIngredient
	Base     UpdateRecipesIngredient
	Addition UpdateRecipesIngredient
	Result   decode.Slot
}

type UpdateRecipesSmithingTrim struct {
	Type     string
	Template UpdateRecipesIngredient
	Base     UpdateRecipesIngredient
	Addition UpdateRecipesIngredient
}

type UpdateRecipes struct {
	Recipes map[string]any
}

func (p UpdateRecipes) Read(reader decode.PacketReader) (packet decode.Packet, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	recipes := map[string]any{}
	for i := int32(0); i < count; i++ {
		rType, err := reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		rId, err := reader.ReadIdentifier()
		if err != nil {
			return nil, err
		}
		switch rType {
		case "minecraft:crafting_shapeless":
			group, err := reader.ReadString()
			if err != nil {
				return nil, err
			}
			cat, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			in, err := readIngredients(reader)
			if err != nil {
				return nil, err
			}
			out, err := reader.ReadSlot()
			if err != nil {
				return nil, err
			}
			recipes[rId] = UpdateRecipesCraftingShapeless{
				Type:        rType,
				Group:       group,
				Category:    cat,
				Ingredients: in,
				Result:      out,
			}
		case "minecraft:crafting_shaped":
			width, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			height, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			group, err := reader.ReadString()
			if err != nil {
				return nil, err
			}
			cat, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			count1 := width * height
			if count1 < 0 {
				return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
			}
			in := make([]UpdateRecipesIngredient, count1)
			for j := int32(0); j < count1; j++ {
				in[j], err = readIngredient(reader)
				if err != nil {
					return nil, err
				}
			}
			out, err := reader.ReadSlot()
			if err != nil {
				return nil, err
			}
			notify, err := reader.ReadBoolean()
			recipes[rId] = UpdateRecipesCraftingShaped{
				Width:            width,
				Height:           height,
				Type:             rType,
				Group:            group,
				Category:         cat,
				Ingredients:      in,
				Result:           out,
				ShowNotification: notify,
			}
		case "minecraft:crafting_special_armordye", "minecraft:crafting_special_bookcloning", "minecraft:crafting_special_mapcloning",
			"minecraft:crafting_special_mapextending", "minecraft:crafting_special_firework_rocket", "minecraft:crafting_special_firework_start",
			"minecraft:crafting_special_firework_star_fade", "minecraft:crafting_special_repairitem", "minecraft:crafting_special_tippedarrow",
			"minecraft:crafting_special_bannerduplicate", "minecraft:crafting_special_shielddecoration", "minecraft:crafting_special_shulkerboxcoloring",
			"minecraft:crafting_special_suspiciousstew", "minecraft:crafting_decorated_pot":
			cat, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			recipes[rId] = UpdateRecipesCraftingSpecial{
				Type:     rType,
				Category: cat,
			}
		case "minecraft:smelting", "minecraft:blasting", "minecraft:smoking", "minecraft:campfire_cooking":
			group, err := reader.ReadString()
			if err != nil {
				return nil, err
			}
			cat, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			ingredient, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			out, err := reader.ReadSlot()
			if err != nil {
				return nil, err
			}
			xp, err := reader.ReadFloat()
			if err != nil {
				return nil, err
			}
			cookTime, err := reader.ReadVarInt()
			if err != nil {
				return nil, err
			}
			recipes[rId] = UpdateRecipesSmelting{
				Type:        rType,
				Group:       group,
				Category:    cat,
				Ingredient:  ingredient,
				Result:      out,
				Experience:  xp,
				CookingTime: cookTime,
			}
		case "minecraft:stonecutting":
			group, err := reader.ReadString()
			if err != nil {
				return nil, err
			}
			ingredient, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			out, err := reader.ReadSlot()
			if err != nil {
				return nil, err
			}
			recipes[rId] = UpdateRecipesCutting{
				Type:       rType,
				Group:      group,
				Ingredient: ingredient,
				Result:     out,
			}
		case "minecraft:smithing_transform":
			template, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			base, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			addition, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			out, err := reader.ReadSlot()
			if err != nil {
				return nil, err
			}
			recipes[rId] = UpdateRecipesSmithingTransform{
				Type:     rType,
				Template: template,
				Base:     base,
				Addition: addition,
				Result:   out,
			}
		case "minecraft:smithing_trim":
			template, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			base, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			addition, err := readIngredient(reader)
			if err != nil {
				return nil, err
			}
			recipes[rId] = UpdateRecipesSmithingTrim{
				Type:     rType,
				Template: template,
				Base:     base,
				Addition: addition,
			}
		}
	}
	return UpdateRecipes{Recipes: recipes}, nil
}

func (p UpdateRecipes) IsValid() (reason error) {
	// TODO: validate?
	return nil
}

func readIngredients(reader decode.PacketReader) (value []UpdateRecipesIngredient, err error) {
	count, err := reader.ReadVarInt()
	if err != nil {
		return nil, err
	} else if count < 0 {
		return nil, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	value = make([]UpdateRecipesIngredient, count)
	for i := int32(0); i < count; i++ {
		value[i], err = readIngredient(reader)
		if err != nil {
			return nil, err
		}
	}
	return value, nil
}

func readIngredient(reader decode.PacketReader) (value UpdateRecipesIngredient, err error) {
	value = UpdateRecipesIngredient{}
	count, err := reader.ReadVarInt()
	if err != nil {
		return value, err
	} else if count < 0 {
		return value, errors.Join(fmt.Errorf("count must be atleast 0 got %d", count), error2.ErrDecodeTooSmall)
	}
	value.Items = make([]decode.Slot, count)
	for i := int32(0); i < count; i++ {
		value.Items[i], err = reader.ReadSlot()
		if err != nil {
			return value, err
		}
	}
	return value, nil
}
