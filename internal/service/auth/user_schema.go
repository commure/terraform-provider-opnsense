package auth

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// userResourceModel describes the resource data model.
type userResourceModel struct {
	Disabled          types.Bool   `tfsdk:"disabled"`
	Name              types.String `tfsdk:"name"`
	Password          types.String `tfsdk:"password"`
	ScrambledPassword types.String `tfsdk:"scrambled_password"`
	Email             types.String `tfsdk:"email"`
	Fullname          types.String `tfsdk:"fullname"`
	Comment           types.String `tfsdk:"comment"`
	OtpSeed           types.String `tfsdk:"otp_seed"`
	AuthorizedKeys    types.String `tfsdk:"authorized_keys"`
	Shell             types.String `tfsdk:"shell"`
	Scope             types.String `tfsdk:"scope"`
	Expires           types.String `tfsdk:"expires"`
	LandingPage       types.String `tfsdk:"landing_page"`
	ApiKeys           types.String `tfsdk:"api_keys"`
	Language          types.String `tfsdk:"language"`
	Dashboard         types.String `tfsdk:"dashboard"`
	Privilege         types.Set    `tfsdk:"privilege"`
	GroupMemberships  types.Set    `tfsdk:"group_memberships"`

	UserId types.String `tfsdk:"user_id"`
	Id     types.String `tfsdk:"id"`
}

func userResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense auth user.",

		Attributes: map[string]schema.Attribute{
			"disabled": schema.BoolAttribute{
				MarkdownDescription: "Disable this user. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The login name of the user.",
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "The user password.",
				Optional:            true,
				Sensitive:           true,
			},
			"scrambled_password": schema.StringAttribute{
				MarkdownDescription: "The scrambled (hashed) password.",
				Optional:            true,
				Sensitive:           true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: "The user email address.",
				Optional:            true,
			},
			"fullname": schema.StringAttribute{
				MarkdownDescription: "The user full name (description).",
				Optional:            true,
			},
			"comment": schema.StringAttribute{
				MarkdownDescription: "Comment for this user.",
				Optional:            true,
			},
			"otp_seed": schema.StringAttribute{
				MarkdownDescription: "OTP seed for two-factor authentication.",
				Optional:            true,
				Sensitive:           true,
			},
			"authorized_keys": schema.StringAttribute{
				MarkdownDescription: "SSH authorized keys.",
				Optional:            true,
			},
			"shell": schema.StringAttribute{
				MarkdownDescription: "The user login shell.",
				Optional:            true,
				Computed:            true,
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: "The user scope.",
				Optional:            true,
				Computed:            true,
			},
			"expires": schema.StringAttribute{
				MarkdownDescription: "Account expiration date.",
				Optional:            true,
			},
			"landing_page": schema.StringAttribute{
				MarkdownDescription: "The landing page after login.",
				Optional:            true,
			},
			"api_keys": schema.StringAttribute{
				MarkdownDescription: "API keys for this user.",
				Optional:            true,
			},
			"language": schema.StringAttribute{
				MarkdownDescription: "The user language preference.",
				Optional:            true,
				Computed:            true,
			},
			"dashboard": schema.StringAttribute{
				MarkdownDescription: "Dashboard columns.",
				Optional:            true,
			},
			"privilege": schema.SetAttribute{
				MarkdownDescription: "Set of privileges assigned to this user. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"group_memberships": schema.SetAttribute{
				MarkdownDescription: "Set of group memberships for this user. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"user_id": schema.StringAttribute{
				MarkdownDescription: "The numeric user ID (UID).",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func userDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense auth user.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"disabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this user is disabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "The login name of the user.",
				Computed:            true,
			},
			"email": dschema.StringAttribute{
				MarkdownDescription: "The user email address.",
				Computed:            true,
			},
			"fullname": dschema.StringAttribute{
				MarkdownDescription: "The user full name (description).",
				Computed:            true,
			},
			"comment": dschema.StringAttribute{
				MarkdownDescription: "Comment for this user.",
				Computed:            true,
			},
			"authorized_keys": dschema.StringAttribute{
				MarkdownDescription: "SSH authorized keys.",
				Computed:            true,
			},
			"shell": dschema.StringAttribute{
				MarkdownDescription: "The user login shell.",
				Computed:            true,
			},
			"scope": dschema.StringAttribute{
				MarkdownDescription: "The user scope.",
				Computed:            true,
			},
			"expires": dschema.StringAttribute{
				MarkdownDescription: "Account expiration date.",
				Computed:            true,
			},
			"landing_page": dschema.StringAttribute{
				MarkdownDescription: "The landing page after login.",
				Computed:            true,
			},
			"api_keys": dschema.StringAttribute{
				MarkdownDescription: "API keys for this user.",
				Computed:            true,
			},
			"language": dschema.StringAttribute{
				MarkdownDescription: "The user language preference.",
				Computed:            true,
			},
			"dashboard": dschema.StringAttribute{
				MarkdownDescription: "Dashboard columns.",
				Computed:            true,
			},
			"privilege": dschema.SetAttribute{
				MarkdownDescription: "Set of privileges assigned to this user.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"group_memberships": dschema.SetAttribute{
				MarkdownDescription: "Set of group memberships for this user.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"user_id": dschema.StringAttribute{
				MarkdownDescription: "The numeric user ID (UID).",
				Computed:            true,
			},
		},
	}
}

func convertUserSchemaToStruct(d *userResourceModel) (*auth.User, error) {
	// Parse 'Privilege'
	var privilegeList []string
	d.Privilege.ElementsAs(context.Background(), &privilegeList, false)

	// Parse 'GroupMemberships'
	var groupMembershipsList []string
	d.GroupMemberships.ElementsAs(context.Background(), &groupMembershipsList, false)

	return &auth.User{
		Name:              d.Name.ValueString(),
		Disabled:          tools.BoolToString(d.Disabled.ValueBool()),
		Password:          d.Password.ValueString(),
		ScrambledPassword: d.ScrambledPassword.ValueString(),
		Email:             d.Email.ValueString(),
		Fullname:          d.Fullname.ValueString(),
		Comment:           d.Comment.ValueString(),
		OtpSeed:           d.OtpSeed.ValueString(),
		AuthorizedKeys:    d.AuthorizedKeys.ValueString(),
		Shell:             api.SelectedMap(d.Shell.ValueString()),
		Scope:             d.Scope.ValueString(),
		Expires:           d.Expires.ValueString(),
		LandingPage:       d.LandingPage.ValueString(),
		ApiKeys:           d.ApiKeys.ValueString(),
		Language:          api.SelectedMap(d.Language.ValueString()),
		Dashboard:         d.Dashboard.ValueString(),
		Privilege:         privilegeList,
		GroupMemberships:  groupMembershipsList,
	}, nil
}

func convertUserStructToSchema(d *auth.User) (*userResourceModel, error) {
	model := &userResourceModel{
		Disabled:          types.BoolValue(tools.StringToBool(d.Disabled)),
		Name:              types.StringValue(d.Name),
		Email:             tools.StringOrNull(d.Email),
		Fullname:          tools.StringOrNull(d.Fullname),
		Comment:           tools.StringOrNull(d.Comment),
		OtpSeed:           tools.StringOrNull(d.OtpSeed),
		AuthorizedKeys:    tools.StringOrNull(d.AuthorizedKeys),
		Shell:             types.StringValue(d.Shell.String()),
		Scope:             types.StringValue(d.Scope),
		Expires:           tools.StringOrNull(d.Expires),
		LandingPage:       tools.StringOrNull(d.LandingPage),
		ApiKeys:           tools.StringOrNull(d.ApiKeys),
		Language:          types.StringValue(d.Language.String()),
		Dashboard:         tools.StringOrNull(d.Dashboard),
		Password:          types.StringNull(),
		ScrambledPassword: types.StringNull(),
		UserId:            tools.StringOrNull(d.UserId),
	}

	// Parse 'Privilege'
	var privilegeList []attr.Value
	for _, i := range d.Privilege {
		if i == "" {
			continue
		}
		privilegeList = append(privilegeList, basetypes.NewStringValue(i))
	}
	privilegeTypeList, _ := types.SetValue(types.StringType, privilegeList)
	model.Privilege = privilegeTypeList

	// Parse 'GroupMemberships'
	var groupMembershipsList []attr.Value
	for _, i := range d.GroupMemberships {
		if i == "" {
			continue
		}
		groupMembershipsList = append(groupMembershipsList, basetypes.NewStringValue(i))
	}
	groupMembershipsTypeList, _ := types.SetValue(types.StringType, groupMembershipsList)
	model.GroupMemberships = groupMembershipsTypeList

	return model, nil
}
