package core

// EntitySchema describes an entity that is used with the endpoints, domain validators and
// serializers.
type EntitySchema struct {
	// Name describes the name of the entity, it must be unique.
	Name string `hcl:"name"`
	// Fields describes fields of given domain entity, i.e. email, username, password etc.
	Fields map[string]any `hcl:"fields"`
}
