package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"terraform-provider-opnsense/internal/tools"
)

// AuthUserResourceModel describes the resource data model.
type AuthUserResourceModel struct {
	UserId            types.String `tfsdk:"user_id"`
	Name              types.String `tfsdk:"name"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Scope             types.String `tfsdk:"scope"`
	Expires           types.String `tfsdk:"expires"`
	AuthorizedKeys    types.String `tfsdk:"authorized_keys"`
	OtpSeed           types.String `tfsdk:"otp_seed"`
	Shell             types.String `tfsdk:"shell"`
	Password          types.String `tfsdk:"password"`
	ScrambledPassword types.String `tfsdk:"scrambled_password"`
	LandingPage       types.String `tfsdk:"landing_page"`
	Comment           types.String `tfsdk:"comment"`
	Email             types.String `tfsdk:"email"`
	ApiKeys           types.String `tfsdk:"api_keys"`
	Privilege         types.Set    `tfsdk:"privilege"`
	Language          types.String `tfsdk:"language"`
	GroupMemberships  types.Set    `tfsdk:"group_memberships"`
	Fullname          types.String `tfsdk:"fullname"`
	Dashboard         types.String `tfsdk:"dashboard"`

	Id types.String `tfsdk:"id"`
}

func AuthUserResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense user account.",

		Attributes: map[string]schema.Attribute{
			"user_id":           schema.StringAttribute{MarkdownDescription: "The numeric user ID.", Optional: true, Computed: true},
			"name":              schema.StringAttribute{MarkdownDescription: "The username.", Required: true},
			"disabled":          schema.BoolAttribute{MarkdownDescription: "Whether the user is disabled.", Optional: true, Computed: true},
			"scope":             schema.StringAttribute{MarkdownDescription: "The user scope.", Optional: true, Computed: true},
			"expires":           schema.StringAttribute{MarkdownDescription: "Expiration date.", Optional: true},
			"authorized_keys":   schema.StringAttribute{MarkdownDescription: "SSH authorized keys.", Optional: true},
			"otp_seed":          schema.StringAttribute{MarkdownDescription: "OTP seed.", Optional: true, Sensitive: true},
			"shell":             schema.StringAttribute{MarkdownDescription: "Login shell.", Optional: true, Computed: true},
			"password":          schema.StringAttribute{MarkdownDescription: "Password.", Optional: true, Sensitive: true},
			"scrambled_password": schema.StringAttribute{MarkdownDescription: "Scrambled password.", Optional: true, Sensitive: true},
			"landing_page":      schema.StringAttribute{MarkdownDescription: "Landing page after login.", Optional: true},
			"comment":           schema.StringAttribute{MarkdownDescription: "Comment.", Optional: true},
			"email":             schema.StringAttribute{MarkdownDescription: "Email address.", Optional: true},
			"api_keys":          schema.StringAttribute{MarkdownDescription: "API keys.", Optional: true},
			"privilege": schema.SetAttribute{
				MarkdownDescription: "Privileges.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"language":    schema.StringAttribute{MarkdownDescription: "Preferred language.", Optional: true, Computed: true},
			"group_memberships": schema.SetAttribute{
				MarkdownDescription: "Group memberships.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"fullname":  schema.StringAttribute{MarkdownDescription: "Full name.", Optional: true},
			"dashboard": schema.StringAttribute{MarkdownDescription: "Dashboard configuration.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func AuthUserDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an OPNsense user account.",

		Attributes: map[string]dschema.Attribute{
			"id":                 dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"user_id":            dschema.StringAttribute{Computed: true, MarkdownDescription: "The numeric user ID."},
			"name":               dschema.StringAttribute{Computed: true, MarkdownDescription: "The username."},
			"disabled":           dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether disabled."},
			"scope":              dschema.StringAttribute{Computed: true, MarkdownDescription: "The user scope."},
			"expires":            dschema.StringAttribute{Computed: true, MarkdownDescription: "Expiration date."},
			"authorized_keys":    dschema.StringAttribute{Computed: true, MarkdownDescription: "SSH authorized keys."},
			"otp_seed":           dschema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "OTP seed."},
			"shell":              dschema.StringAttribute{Computed: true, MarkdownDescription: "Login shell."},
			"password":           dschema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Password."},
			"scrambled_password": dschema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Scrambled password."},
			"landing_page":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Landing page."},
			"comment":            dschema.StringAttribute{Computed: true, MarkdownDescription: "Comment."},
			"email":              dschema.StringAttribute{Computed: true, MarkdownDescription: "Email."},
			"api_keys":           dschema.StringAttribute{Computed: true, MarkdownDescription: "API keys."},
			"privilege":          dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Privileges."},
			"language":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Language."},
			"group_memberships":  dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Group memberships."},
			"fullname":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Full name."},
			"dashboard":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Dashboard."},
		},
	}
}

func convertAuthUserSchemaToStruct(d *AuthUserResourceModel) (*auth.User, error) {
	var privList []string
	d.Privilege.ElementsAs(context.Background(), &privList, false)
	var groupList []string
	d.GroupMemberships.ElementsAs(context.Background(), &groupList, false)

	return &auth.User{
		UserId:            d.UserId.ValueString(),
		Name:              d.Name.ValueString(),
		Disabled:          tools.BoolToString(d.Disabled.ValueBool()),
		Scope:             d.Scope.ValueString(),
		Expires:           d.Expires.ValueString(),
		AuthorizedKeys:    d.AuthorizedKeys.ValueString(),
		OtpSeed:           d.OtpSeed.ValueString(),
		Shell:             api.SelectedMap(d.Shell.ValueString()),
		Password:          d.Password.ValueString(),
		ScrambledPassword: d.ScrambledPassword.ValueString(),
		LandingPage:       d.LandingPage.ValueString(),
		Comment:           d.Comment.ValueString(),
		Email:             d.Email.ValueString(),
		ApiKeys:           d.ApiKeys.ValueString(),
		Privilege:         privList,
		Language:          api.SelectedMap(d.Language.ValueString()),
		GroupMemberships:  groupList,
		Fullname:          d.Fullname.ValueString(),
		Dashboard:         d.Dashboard.ValueString(),
	}, nil
}

func convertAuthUserStructToSchema(d *auth.User) (*AuthUserResourceModel, error) {
	model := &AuthUserResourceModel{
		UserId:            tools.StringOrNull(d.UserId),
		Name:              types.StringValue(d.Name),
		Disabled:          types.BoolValue(tools.StringToBool(d.Disabled)),
		Scope:             tools.StringOrNull(d.Scope),
		Expires:           tools.StringOrNull(d.Expires),
		AuthorizedKeys:    tools.StringOrNull(d.AuthorizedKeys),
		OtpSeed:           tools.StringOrNull(d.OtpSeed),
		Shell:             types.StringValue(d.Shell.String()),
		Password:          tools.StringOrNull(d.Password),
		ScrambledPassword: tools.StringOrNull(d.ScrambledPassword),
		LandingPage:       tools.StringOrNull(d.LandingPage),
		Comment:           tools.StringOrNull(d.Comment),
		Email:             tools.StringOrNull(d.Email),
		ApiKeys:           tools.StringOrNull(d.ApiKeys),
		Language:          types.StringValue(d.Language.String()),
		Fullname:          tools.StringOrNull(d.Fullname),
		Dashboard:         tools.StringOrNull(d.Dashboard),
	}

	setFromSlice := func(s []string) types.Set {
		var list []attr.Value
		for _, i := range s {
			if i == "" { continue }
			list = append(list, basetypes.NewStringValue(i))
		}
		v, _ := types.SetValue(types.StringType, list)
		return v
	}

	model.Privilege = setFromSlice(d.Privilege)
	model.GroupMemberships = setFromSlice(d.GroupMemberships)

	return model, nil
}
