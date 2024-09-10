1] Generics 
- Generics allow you to write functions, types, and data structures that can operate on different data types while maintaining type safety.
- The ability to write reusable code that can work with various types without sacrificing type safety.

Key Concepts:
1]Type Parameters: Functions or types can accept type parameters, which are placeholders for specific types
2]Type Constraints :Type constraints define the acceptable types that can be passed as type parameters

Different type of Type Constraints
1. The any Constraint(The [any] constraint (or the empty interface interface{}).)
The [any] constraint is a special type constraint that allows any type to be used. It’s shorthand for the empty interface interface{}.


In Go, type constraints specify the permissible types for generic type parameters, allowing you to define how a generic function or type behaves with different sets of types. Constraints can be based on:

The any constraint (or the empty interface interface{}).
Built-in interface constraints like comparable or `constraints from the standard library**.
Custom interface constraints defined by you.
Union of specific types (type sets).
Type parameters with methods.
1. The any Constraint
The any constraint is a special type constraint that allows any type to be used. It’s shorthand for the empty interface interface{}.

2. The comparable Constraint
The comparable constraint restricts the type to those that can be compared using the == and != operators. 
Can be compared, like int, string, or structs where all fields are comparable.
This is useful for functions that need to compare values, like implementing a search or equality check.

3. Built-in Interface Constraints
Go's standard library provides several built-in interfaces that can be used as type constraints.
The constraints.Ordered interface allows any type that supports <, >, <=, >=, like numeric types and strings.

4. Custom Interface Constraints
You can define custom interfaces to constrain generic types based on method sets. The type must implement the methods of the interface to be used as a type parameter.

5. Union of Specific Types (Type Sets)
You can define a constraint that allows a set of specific types using a type union. This is useful when you want to limit the generic function to only certain types.
For example T is constrained to be either int or float64. No other type is allowed.

Summary of Common Type Constraints:
any: Allows any type.
comparable: Allows types that can be compared with == and !=.
Built-in interfaces: Like constraints.Ordered for ordered types.
Custom interfaces: Defined by you, specifying methods the type must implement.
Union of types: Specifies a set of types using type sets.
Methods as constraints: Specify methods that types must implement.