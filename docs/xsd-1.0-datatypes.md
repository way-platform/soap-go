# XML Schema Part 2: Datatypes

Original: https://www.w3.org/TR/xmlschema-2/

## Abstract

_XML Schema: Datatypes_ is part 2 of the specification of the XML Schema language. It defines facilities for defining datatypes to be used in XML Schemas as well as other XML specifications. The datatype language, which is itself represented in XML 1.0, provides a superset of the capabilities found in XML 1.0 document type definitions (DTDs) for specifying datatypes on elements and attributes.

## 1 Introduction

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#requirements)1.1 Purpose

The [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML) specification defines limited facilities for applying datatypes to document content in that documents may contain or refer to DTDs that assign types to elements and attributes. However, document authors, including authors of traditional _documents_ and those transporting _data_ in XML, often require a higher degree of type checking to ensure robustness in document understanding and data interchange.

The table below offers two typical examples of XML instances in which datatypes are implicit: the instance on the left represents a billing invoice, the instance on the right a memo or perhaps an email message in XML.

|Data oriented|Document oriented|
|---|---|
|<invoice><br>  <orderDate>1999-01-21</orderDate><br>  <shipDate>1999-01-25</shipDate><br>  <billingAddress><br>   <name>Ashok Malhotra</name><br>   <street>123 Microsoft Ave.</street><br>   <city>Hawthorne</city><br>   <state>NY</state><br>   <zip>10532-0000</zip><br>  </billingAddress><br>  <voice>555-1234</voice><br>  <fax>555-4321</fax><br></invoice>|<memo importance='high'<br>      date='1999-03-23'><br>  <from>Paul V. Biron</from><br>  <to>Ashok Malhotra</to><br>  <subject>Latest draft</subject><br>  <body><br>    We need to discuss the latest<br>    draft <emph>immediately</emph>.<br>    Either email me at <email><br>    mailto:paul.v.biron@kp.org</email><br>    or call <phone>555-9876</phone><br>  </body><br></memo>|

The invoice contains several dates and telephone numbers, the postal abbreviation for a state (which comes from an enumerated list of sanctioned values), and a ZIP code (which takes a definable regular form). The memo contains many of the same types of information: a date, telephone number, email address and an "importance" value (from an enumerated list, such as "low", "medium" or "high"). Applications which process invoices and memos need to raise exceptions if something that was supposed to be a date or telephone number does not conform to the rules for valid dates or telephone numbers.

In both cases, validity constraints exist on the content of the instances that are not expressible in XML DTDs. The limited datatyping facilities in XML have prevented validating XML processors from supplying the rigorous type checking required in these situations. The result has been that individual applications writers have had to implement type checking in an ad hoc manner. This specification addresses the need of both document authors and applications writers for a robust, extensible datatype system for XML which could be incorporated into XML processors. As discussed below, these datatypes could be used in other XML-related standards as well.

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#purpose) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#scope)1.2 Requirements

The [[XML Schema Requirements]](https://www.w3.org/TR/xmlschema-2/#schema-requirements) document spells out concrete requirements to be fulfilled by this specification, which state that the XML Schema Language must:

1. provide for primitive data typing, including byte, date, integer, sequence, SQL and Java primitive datatypes, etc.;
2. define a type system that is adequate for import/export from database systems (e.g., relational, object, OLAP);
3. distinguish requirements relating to lexical data representation vs. those governing an underlying information set;
4. allow creation of user-defined datatypes, such as datatypes that are derived from existing datatypes and which may constrain certain of its properties (e.g., range, precision, length, format).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#requirements) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#terminology)1.3 Scope

This portion of the XML Schema Language discusses datatypes that can be used in an XML Schema. These datatypes can be specified for element content that would be specified as [#PCDATA](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-chardata) and attribute values of [various types](https://www.w3.org/TR/2000/WD-xml-2e-20000814#sec-attribute-types) in a DTD. It is the intention of this specification that it be usable outside of the context of XML Schemas for a wide range of other XML-related activities such as [[XSL]](https://www.w3.org/TR/xmlschema-2/#XSL) and [[RDF Schema]](https://www.w3.org/TR/xmlschema-2/#RDFSchema).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#scope) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#constraints-and-contributions)1.4 Terminology

The terminology used to describe XML Schema Datatypes is defined in the body of this specification. The terms defined in the following list are used in building those definitions and in describing the actions of a datatype processor:

[Definition:]   for compatibility

A feature of this specification included solely to ensure that schemas which use this feature remain compatible with [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML)

[Definition:]  **may**

Conforming documents and processors are permitted to but need not behave as described.

[Definition:]  **match**

(Of strings or names:) Two strings or names being compared must be identical. Characters with multiple possible representations in ISO/IEC 10646 (e.g. characters with both precomposed and base+diacritic forms) match only if they have the same representation in both strings. No case folding is performed. (Of strings and rules in the grammar:) A string matches a grammatical production if it belongs to the language generated by that production.

[Definition:]  **must**

Conforming documents and processors are required to behave as described; otherwise they are in [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error).

[Definition:]  **error**

A violation of the rules of this specification; results are undefined. Conforming software [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) detect and report an **error** and [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) recover from it.

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#terminology)1.5 Constraints and Contributions

This specification provides three different kinds of normative statements about schema components, their representations in XML and their contribution to the schema-validation of information items:

[Definition:]   **Constraint on Schemas**

Constraints on the schema components themselves, i.e. conditions components [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) satisfy to be components at all. Largely to be found in [Datatype components (§4)](https://www.w3.org/TR/xmlschema-2/#datatype-components).

[Definition:]   **Schema Representation Constraint**

Constraints on the representation of schema components in XML. Some but not all of these are expressed in [Schema for Datatype Definitions (normative) (§A)](https://www.w3.org/TR/xmlschema-2/#schema) and [DTD for Datatype Definitions (non-normative) (§B)](https://www.w3.org/TR/xmlschema-2/#dtd-for-datatypeDefs).

[Definition:]   **Validation Rule**

Constraints expressed by schema components which information items [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) satisfy to be schema-valid. Largely to be found in [Datatype components (§4)](https://www.w3.org/TR/xmlschema-2/#datatype-components).

## 2 Type System

This section describes the conceptual framework behind the type system defined in this specification. The framework has been influenced by the [[ISO 11404]](https://www.w3.org/TR/xmlschema-2/#ISO11404) standard on language-independent datatypes as well as the datatypes for [[SQL]](https://www.w3.org/TR/xmlschema-2/#SQL) and for programming languages such as Java.

The datatypes discussed in this specification are computer representations of well known abstract concepts such as _integer_ and _date_. It is not the place of this specification to define these abstract concepts; many other publications provide excellent definitions.

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#value-space)2.1 Datatype

[Definition:]  In this specification, a **datatype** is a 3-tuple, consisting of a) a set of distinct values, called its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), b) a set of lexical representations, called its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space), and c) a set of [·facet·](https://www.w3.org/TR/xmlschema-2/#dt-facet)s that characterize properties of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), individual values or lexical items.

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#datatype) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#lexical-space)2.2 Value space

[Definition:]  A **value space** is the set of values for a given datatype. Each value in the **value space** of a datatype is denoted by one or more literals in its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space).

The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of a given datatype can be defined in one of the following ways:

- defined axiomatically from fundamental notions (intensional definition) [see [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive)]
- enumerated outright (extensional definition) [see [·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration)]
- defined by restricting the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of an already defined datatype to a particular subset with a given set of properties [see [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived)]
- defined as a combination of values from one or more already defined [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)(s) by a specific construction procedure [see [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) and [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union)]

[·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s have certain properties. For example, they always have the property of [·cardinality·](https://www.w3.org/TR/xmlschema-2/#dt-cardinality), some definition of _equality_ and might be [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered), by which individual values within the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) can be compared to one another. The properties of [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s that are recognized by this specification are defined in [Fundamental facets (§2.4.1)](https://www.w3.org/TR/xmlschema-2/#fundamental-facets).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#value-space) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#facets)2.3 Lexical space

In addition to its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), each datatype also has a lexical space.

[Definition:]  A **lexical space** is the set of valid _literals_ for a datatype.

For example, "100" and "1.0E2" are two different literals from the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of [float](https://www.w3.org/TR/xmlschema-2/#float) which both denote the same value. The type system defined in this specification provides a mechanism for schema designers to control the set of values and the corresponding set of acceptable literals of those values for a datatype.

**Note:**  The literals in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s defined in this specification have the following characteristics:

Interoperability:

The number of literals for each value has been kept small; for many datatypes there is a one-to-one mapping between literals and values. This makes it easy to exchange the values between different systems. In many cases, conversion from locale-dependent representations will be required on both the originator and the recipient side, both for computer processing and for interaction with humans.

Basic readability:

Textual, rather than binary, literals are used. This makes hand editing, debugging, and similar activities possible.

Ease of parsing and serializing:

Where possible, literals correspond to those found in common programming languages and libraries.

#### 2.3.1 Canonical Lexical Representation

While the datatypes defined in this specification have, for the most part, a single lexical representation i.e. each value in the datatype's [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is denoted by a single literal in its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space), this is not always the case. The example in the previous section showed two literals for the datatype [float](https://www.w3.org/TR/xmlschema-2/#float) which denote the same value. Similarly, there [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) be several literals for one of the date or time datatypes that denote the same value using different timezone indicators.

[Definition:]  A **canonical lexical representation** is a set of literals from among the valid set of literals for a datatype such that there is a one-to-one mapping between literals in the **canonical lexical representation** and values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#lexical-space) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#datatype-dichotomies)2.4 Facets

        2.4.1 [Fundamental facets](https://www.w3.org/TR/xmlschema-2/#fundamental-facets)
        2.4.2 [Constraining or Non-fundamental facets](https://www.w3.org/TR/xmlschema-2/#non-fundamental)

[Definition:]  A **facet** is a single defining aspect of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space). Generally speaking, each facet characterizes a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) along independent axes or dimensions.

The facets of a datatype serve to distinguish those aspects of one datatype which _differ_ from other datatypes. Rather than being defined solely in terms of a prose description the datatypes in this specification are defined in terms of the _synthesis_ of facet values which together determine the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and properties of the datatype.

Facets are of two types: _fundamental_ facets that define the datatype and _non-fundamental_ or _constraining_ facets that constrain the permitted values of a datatype.

#### 2.4.1 Fundamental facets

[Definition:]   A **fundamental facet** is an abstract property which serves to semantically characterize the values in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

All **fundamental facets** are fully described in [Fundamental Facets (§4.2)](https://www.w3.org/TR/xmlschema-2/#rf-fund-facets).

#### 2.4.2 Constraining or Non-fundamental facets

[Definition:]  A **constraining facet** is an optional property that can be applied to a datatype to constrain its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

Constraining the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) consequently constrains the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space). Adding [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s to a [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) is described in [Derivation by restriction (§4.1.2.1)](https://www.w3.org/TR/xmlschema-2/#derivation-by-restriction).

All **constraining facets** are fully described in [Constraining Facets (§4.3)](https://www.w3.org/TR/xmlschema-2/#rf-facets).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#facets)2.5 Datatype dichotomies

        2.5.1 [Atomic vs. list vs. union datatypes](https://www.w3.org/TR/xmlschema-2/#atomic-vs-list)
        2.5.2 [Primitive vs. derived datatypes](https://www.w3.org/TR/xmlschema-2/#primitive-vs-derived)
        2.5.3 [Built-in vs. user-derived datatypes](https://www.w3.org/TR/xmlschema-2/#built-in-vs-user-derived)

It is useful to categorize the datatypes defined in this specification along various dimensions, forming a set of characterization dichotomies.

#### 2.5.1 Atomic vs. list vs. union datatypes

The first distinction to be made is that between [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic), [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) and [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatypes.

- [Definition:]  **Atomic** datatypes are those having values which are regarded by this specification as being indivisible.
- [Definition:]  **List** datatypes are those having values each of which consists of a finite-length (possibly empty) sequence of values of an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype.
- [Definition:]  **Union** datatypes are those whose [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s are the union of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s of one or more other datatypes.

For example, a single token which [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match)es [Nmtoken](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Nmtoken) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML) could be the value of an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype ([NMTOKEN](https://www.w3.org/TR/xmlschema-2/#NMTOKEN)); while a sequence of such tokens could be the value of a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype ([NMTOKENS](https://www.w3.org/TR/xmlschema-2/#NMTOKENS)).

##### 2.5.1.1 Atomic datatypes

[·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatypes can be either [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) or [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype is a set of "atomic" values, which for the purposes of this specification, are not further decomposable. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype is a set of _literals_ whose internal structure is specific to the datatype in question.

##### 2.5.1.2 List datatypes

Several type systems (such as the one described in [[ISO 11404]](https://www.w3.org/TR/xmlschema-2/#ISO11404)) treat [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatypes as special cases of the more general notions of aggregate or collection datatypes.

[·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatypes are always [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype is a set of finite-length sequences of [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) values. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype is a set of literals whose internal structure is a space-separated sequence of literals of the [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype of the items in the [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list).

[Definition:]   The [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype that participates in the definition of a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype is known as the **itemType** of that [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype.

Example

<simpleType name='sizes'>
  <list itemType='decimal'/>
</simpleType>

<cerealSizes xsi:type='sizes'> 8 10.5 12 </cerealSizes>

A [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype can be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype whose [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) allows space (such as [string](https://www.w3.org/TR/xmlschema-2/#string) or [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI))or a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype any of whose [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes)'s [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) allows space. In such a case, regardless of the input, list items will be separated at space boundaries.

Example

<simpleType name='listOfString'>
  <list itemType='string'/>
</simpleType>

<someElement xsi:type='listOfString'>
this is not list item 1
this is not list item 2
this is not list item 3
</someElement>

In the above example, the value of the _someElement_ element is not a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length) 3; rather, it is a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length) 18.

When a datatype is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype, the following [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s apply:

- [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length)
- [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength)
- [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength)
- [·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration)
- [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern)
- [·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace)

For each of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length), [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) and [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength), the _unit of length_ is measured in number of list items. The value of [·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace) is fixed to the value _collapse_.

For [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatypes the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) is composed of space-separated literals of its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType). Hence, any [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) specified when a new datatype is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype is matched against each literal of the [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype and not against the literals of the datatype that serves as its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType).

Example

<xs:simpleType name='myList'>
	<xs:list itemType='xs:integer'/>
</xs:simpleType>
<xs:simpleType name='myRestrictedList'>
	<xs:restriction base='myList'>
		<xs:pattern value='123 (\d+\s)*456'/>
	</xs:restriction>
</xs:simpleType>
<someElement xsi:type='myRestrictedList'>123 456</someElement>
<someElement xsi:type='myRestrictedList'>123 987 456</someElement>
<someElement xsi:type='myRestrictedList'>123 987 567 456</someElement>

The [canonical-lexical-representation](https://www.w3.org/TR/xmlschema-2/#canonical-lexical-representation) for the [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype is defined as the lexical form in which each item in the [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) has the canonical lexical representation of its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType).

##### 2.5.1.3 Union datatypes

The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype are the union of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s of its [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes). [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatypes are always [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived). Currently, there are no [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatypes.

Example

A prototypical example of a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) type is the [maxOccurs attribute](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#p-max_occurs) on the [element element](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-element) in XML Schema itself: it is a union of nonNegativeInteger and an enumeration with the single member, the string "unbounded", as shown below.

  <attributeGroup name="occurs">
    <attribute name="minOccurs" type="nonNegativeInteger"
    	use="optional" default="1"/>
    <attribute name="maxOccurs"use="optional" default="1">
      <simpleType>
        <union>
          <simpleType>
            <restriction base='nonNegativeInteger'/>
          </simpleType>
          <simpleType>
            <restriction base='string'>
              <enumeration value='unbounded'/>
            </restriction>
          </simpleType>
        </union>
      </simpleType>
    </attribute>
  </attributeGroup>

Any number (greater than 1) of [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) [·datatype·](https://www.w3.org/TR/xmlschema-2/#dt-datatype)s can participate in a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) type.

[Definition:]   The datatypes that participate in the definition of a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype are known as the **memberTypes** of that [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype.

The order in which the [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes) are specified in the definition (that is, the order of the <simpleType> children of the <union> element, or the order of the [QName](https://www.w3.org/TR/xmlschema-2/#QName)s in the _memberTypes_ attribute) is significant. During validation, an element or attribute's value is validated against the [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes) in the order in which they appear in the definition until a match is found. The evaluation order can be overridden with the use of [xsi:type](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#xsi_type).

Example

For example, given the definition below, the first instance of the <size> element validates correctly as an [integer (§3.3.13)](https://www.w3.org/TR/xmlschema-2/#integer), the second and third as [string (§3.2.1)](https://www.w3.org/TR/xmlschema-2/#string).

  <xsd:element name='size'>
    <xsd:simpleType>
      <xsd:union>
        <xsd:simpleType>
          <xsd:restriction base='integer'/>
        </xsd:simpleType>
        <xsd:simpleType>
          <xsd:restriction base='string'/>
        </xsd:simpleType>
      </xsd:union>
    </xsd:simpleType>
  </xsd:element>

  <size>1</size>
  <size>large</size>
  <size xsi:type='xsd:string'>1</size>

The [canonical-lexical-representation](https://www.w3.org/TR/xmlschema-2/#canonical-lexical-representation) for a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype is defined as the lexical form in which the values have the canonical lexical representation of the appropriate [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes).

**Note:**  A datatype which is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) in this specification need not be an "atomic" datatype in any programming language used to implement this specification. Likewise, a datatype which is a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) in this specification need not be a "list" datatype in any programming language used to implement this specification. Furthermore, a datatype which is a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) in this specification need not be a "union" datatype in any programming language used to implement this specification.

#### 2.5.2 Primitive vs. derived datatypes

Next, we distinguish between [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) and [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes.

- [Definition:]  **Primitive** datatypes are those that are not defined in terms of other datatypes; they exist _ab initio_.
- [Definition:]  **Derived** datatypes are those that are defined in terms of other datatypes.

For example, in this specification, [float](https://www.w3.org/TR/xmlschema-2/#float) is a well-defined mathematical concept that cannot be defined in terms of other datatypes, while a [integer](https://www.w3.org/TR/xmlschema-2/#integer) is a special case of the more general datatype [decimal](https://www.w3.org/TR/xmlschema-2/#decimal).

[Definition:]   The simple ur-type definition is a special restriction of the [ur-type definition](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-urType) whose name is **anySimpleType** in the XML Schema namespace. **anySimpleType** can be considered as the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of all [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes. **anySimpleType** is considered to have an unconstrained lexical space and a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) consisting of the union of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of all the [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes and the set of all lists of all members of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of all the [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes.

The datatypes defined by this specification fall into both the [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) and [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) categories. It is felt that a judiciously chosen set of [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes will serve the widest possible audience by providing a set of convenient datatypes that can be used as is, as well as providing a rich enough base from which the variety of datatypes needed by schema designers can be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived).

In the example above, [integer](https://www.w3.org/TR/xmlschema-2/#integer) is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [decimal](https://www.w3.org/TR/xmlschema-2/#decimal).

**Note:**  A datatype which is [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) in this specification need not be a "primitive" datatype in any programming language used to implement this specification. Likewise, a datatype which is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) in this specification need not be a "derived" datatype in any programming language used to implement this specification.

As described in more detail in [XML Representation of Simple Type Definition Schema Components (§4.1.2)](https://www.w3.org/TR/xmlschema-2/#xr-defn), each [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be defined in terms of another datatype in one of three ways: 1) by assigning [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s which serve to _restrict_ the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype to a subset of that of the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype); 2) by creating a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype whose [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) consists of finite-length sequences of values of its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType); or 3) by creating a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype whose [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) consists of the union of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of its [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes).

##### 2.5.2.1 Derived by restriction

[Definition:]  A datatype is said to be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by **restriction** from another datatype when values for zero or more [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s are specified that serve to constrain its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and/or its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) to a subset of those of its [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype).

[Definition:]  Every datatype that is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by **restriction** is defined in terms of an existing datatype, referred to as its **base type**. **base type**s can be either [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) or [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived).

##### 2.5.2.2 Derived by list

A [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype can be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from another datatype (its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType)) by creating a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) that consists of a finite-length sequence of values of its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType).

##### 2.5.2.3 Derived by union

One datatype can be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from one or more datatypes by [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union)ing their [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s and, consequently, their [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s.

#### 2.5.3 Built-in vs. user-derived datatypes

- [Definition:]  **Built-in** datatypes are those which are defined in this specification, and can be either [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) or [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived);
- [Definition:]   **User-derived** datatypes are those [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes that are defined by individual schema designers.

Conceptually there is no difference between the [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes included in this specification and the [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatypes which will be created by individual schema designers. The [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes are those which are believed to be so common that if they were not defined in this specification many schema designers would end up "reinventing" them. Furthermore, including these [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes in this specification serves to demonstrate the mechanics and utility of the datatype generation facilities of this specification.

**Note:**  A datatype which is [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) in this specification need not be a "built-in" datatype in any programming language used to implement this specification. Likewise, a datatype which is [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) in this specification need not be a "user-derived" datatype in any programming language used to implement this specification.

## 3 Built-in datatypes

![Diagram of built-in type hierarchy](https://www.w3.org/TR/xmlschema-2/type-hierarchy.gif)

Each built-in datatype in this specification (both [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) and [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived)) can be uniquely addressed via a URI Reference constructed as follows:

1. the base URI is the URI of the XML Schema namespace
2. the fragment identifier is the name of the datatype

For example, to address the [int](https://www.w3.org/TR/xmlschema-2/#int) datatype, the URI is:

- `http://www.w3.org/2001/XMLSchema#int`

Additionally, each facet definition element can be uniquely addressed via a URI constructed as follows:

1. the base URI is the URI of the XML Schema namespace
2. the fragment identifier is the name of the facet

For example, to address the maxInclusive facet, the URI is:

- `http://www.w3.org/2001/XMLSchema#maxInclusive`

Additionally, each facet usage in a built-in datatype definition can be uniquely addressed via a URI constructed as follows:

1. the base URI is the URI of the XML Schema namespace
2. the fragment identifier is the name of the datatype, followed by a period (".") followed by the name of the facet

For example, to address the usage of the maxInclusive facet in the definition of int, the URI is:

- `http://www.w3.org/2001/XMLSchema#int.maxInclusive`

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#built-in-primitive-datatypes)3.1 Namespace considerations

The [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes defined by this specification are designed to be used with the XML Schema definition language as well as other XML specifications. To facilitate usage within the XML Schema definition language, the [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes in this specification have the namespace name:

- http://www.w3.org/2001/XMLSchema

To facilitate usage in specifications other than the XML Schema definition language, such as those that do not want to know anything about aspects of the XML Schema definition language other than the datatypes, each [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatype is also defined in the namespace whose URI is:

- http://www.w3.org/2001/XMLSchema-datatypes

This applies to both [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) and [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes.

Each [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype is also associated with a unique namespace. However, [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatypes do not come from the namespace defined by this specification; rather, they come from the namespace of the schema in which they are defined (see [XML Representation of Schemas](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#declare-schema) in [[XML Schema Part 1: Structures]](https://www.w3.org/TR/xmlschema-2/#structural-schemas)).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#namespaces) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#built-in-derived)3.2 Primitive datatypes

        3.2.1 [string](https://www.w3.org/TR/xmlschema-2/#string)
        3.2.2 [boolean](https://www.w3.org/TR/xmlschema-2/#boolean)
        3.2.3 [decimal](https://www.w3.org/TR/xmlschema-2/#decimal)
        3.2.4 [float](https://www.w3.org/TR/xmlschema-2/#float)
        3.2.5 [double](https://www.w3.org/TR/xmlschema-2/#double)
        3.2.6 [duration](https://www.w3.org/TR/xmlschema-2/#duration)
        3.2.7 [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime)
        3.2.8 [time](https://www.w3.org/TR/xmlschema-2/#time)
        3.2.9 [date](https://www.w3.org/TR/xmlschema-2/#date)
        3.2.10 [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth)
        3.2.11 [gYear](https://www.w3.org/TR/xmlschema-2/#gYear)
        3.2.12 [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay)
        3.2.13 [gDay](https://www.w3.org/TR/xmlschema-2/#gDay)
        3.2.14 [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth)
        3.2.15 [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary)
        3.2.16 [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary)
        3.2.17 [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI)
        3.2.18 [QName](https://www.w3.org/TR/xmlschema-2/#QName)
        3.2.19 [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION)

The [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes defined by this specification are described below. For each datatype, the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) are defined, [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s which apply to the datatype are listed and any datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from this datatype are specified.

[·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes can only be added by revisions to this specification.

#### 3.2.1 string

[Definition:]  The **string** datatype represents character strings in XML. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **string** is the set of finite-length sequences of [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s (as defined in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML)) that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [Char](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Char) production from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). A [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character) is an atomic unit of communication; it is not further specified except to note that every [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character) has a corresponding Universal Character Set code point, which is an integer.

**Note:**  Many human languages have writing systems that require child elements for control of aspects such as bidirectional formating or ruby annotation (see [[Ruby]](https://www.w3.org/TR/xmlschema-2/#ruby) and Section 8.2.4 [Overriding the bidirectional algorithm: the BDO element](https://www.w3.org/TR/1999/REC-html401-19991224/struct/dirlang.html#h-8.2.4) of [[HTML 4.01]](https://www.w3.org/TR/xmlschema-2/#html4)). Thus, **string**, as a simple type that can contain only characters but not child elements, is often not suitable for representing text. In such situations, a complex type that allows mixed content should be considered. For more information, see Section 5.5 [Any Element, Any Attribute](https://www.w3.org/TR/2001/REC-xmlschema-0-20010502/#textType) of [[XML Schema Language: Part 0 Primer]](https://www.w3.org/TR/xmlschema-2/#schema-primer).

**Note:**  As noted in [ordered](https://www.w3.org/TR/xmlschema-2/#dc-ordered), the fact that this specification does not specify an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) for [·string·](https://www.w3.org/TR/xmlschema-2/#dt-string) does not preclude other applications from treating strings as being ordered.

##### 3.2.1.1 Constraining facets

**string** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.2.1.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **string**:

- [normalizedString](https://www.w3.org/TR/xmlschema-2/#normalizedString)

#### 3.2.2 boolean

[Definition:]  **boolean** has the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) required to support the mathematical concept of binary-valued logic: {true, false}.

##### 3.2.2.1 Lexical representation

An instance of a datatype that is defined as [·boolean·](https://www.w3.org/TR/xmlschema-2/#dt-boolean) can have the following legal literals {true, false, 1, 0}.

##### 3.2.2.2 Canonical representation

The canonical representation for **boolean** is the set of literals {true, false}.

##### 3.2.2.3 Constraining facets

**boolean** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

#### 3.2.3 decimal

[Definition:]  **decimal** represents a subset of the real numbers, which can be represented by decimal numerals. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **decimal** is the set of numbers that can be obtained by multiplying an integer by a non-positive power of ten, i.e., expressible as _i × 10^-n_ where _i_ and _n_ are integers and _n >= 0_. Precision is not reflected in this value space; the number 2.0 is not distinct from the number 2.00. The [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) on **decimal** is the order relation on real numbers, restricted to this subset.

**Note:**  All [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) support decimal numbers with a minimum of 18 decimal digits (i.e., with a [·totalDigits·](https://www.w3.org/TR/xmlschema-2/#dt-totalDigits) of 18). However, [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) set an application-defined limit on the maximum number of decimal digits they are prepared to support, in which case that application-defined maximum number [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be clearly documented.

##### 3.2.3.1 Lexical representation

**decimal** has a lexical representation consisting of a finite-length sequence of decimal digits (#x30-#x39) separated by a period as a decimal indicator. An optional leading sign is allowed. If the sign is omitted, "+" is assumed. Leading and trailing zeroes are optional. If the fractional part is zero, the period and following zero(es) can be omitted. For example: `-1.23, 12678967.543233, +100000.00, 210`.

##### 3.2.3.2 Canonical representation

The canonical representation for **decimal** is defined by prohibiting certain options from the [Lexical representation (§3.2.3.1)](https://www.w3.org/TR/xmlschema-2/#decimal-lexical-representation). Specifically, the preceding optional "+" sign is prohibited. The decimal point is required. Leading and trailing zeroes are prohibited subject to the following: there must be at least one digit to the right and to the left of the decimal point which may be a zero.

##### 3.2.3.3 Constraining facets

**decimal** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.2.3.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **decimal**:

- [integer](https://www.w3.org/TR/xmlschema-2/#integer)

#### 3.2.4 float

[Definition:]  **float** is patterned after the IEEE single-precision 32-bit floating point type [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754). The basic [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **float** consists of the values _m × 2^e_, where _m_ is an integer whose absolute value is less than _2^24_, and _e_ is an integer between -149 and 104, inclusive. In addition to the basic [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) described above, the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **float** also contains the following three _special values_: positive and negative infinity and not-a-number (NaN). The [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) on **float** is: _x < y iff y - x_ is positive for x and y in the value space. Positive infinity is greater than all other non-NaN values. NaN equals itself but is [·incomparable·](https://www.w3.org/TR/xmlschema-2/#dt-incomparable) with (neither greater than nor less than) any other value in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

**Note:**  "Equality" in this Recommendation is defined to be "identity" (i.e., values that are identical in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) are equal and vice versa). Identity must be used for the few operations that are defined in this Recommendation. Applications using any of the datatypes defined in this Recommendation may use different definitions of equality for computational purposes; [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754)-based computation systems are examples. Nothing in this Recommendation should be construed as requiring that such applications use identity as their equality relationship when computing.

Any value [·incomparable·](https://www.w3.org/TR/xmlschema-2/#dt-incomparable) with the value used for the four bounding facets ([·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive), [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive), [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive), and [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive)) will be excluded from the resulting restricted [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space). In particular, when "NaN" is used as a facet value for a bounding facet, since no other **float** values are [·comparable·](https://www.w3.org/TR/xmlschema-2/#dt-comparable) with it, the result is a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) either having NaN as its only member (the inclusive cases) or that is empty (the exclusive cases). If any other value is used for a bounding facet, NaN will be excluded from the resulting restricted [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space); to add NaN back in requires union with the NaN-only space.

This datatype differs from that of [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754) in that there is only one NaN and only one zero. This makes the equality and ordering of values in the data space differ from that of [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754) only in that for schema purposes NaN = NaN.

A literal in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) representing a decimal number _d_ maps to the normalized value in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **float** that is closest to _d_ in the sense defined by [[Clinger, WD (1990)]](https://www.w3.org/TR/xmlschema-2/#clinger1990); if _d_ is exactly halfway between two such values then the even value is chosen.

##### 3.2.4.1 Lexical representation

**float** values have a lexical representation consisting of a mantissa followed, optionally, by the character "E" or "e", followed by an exponent. The exponent [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be an [integer](https://www.w3.org/TR/xmlschema-2/#integer). The mantissa must be a [decimal](https://www.w3.org/TR/xmlschema-2/#decimal) number. The representations for exponent and mantissa must follow the lexical rules for [integer](https://www.w3.org/TR/xmlschema-2/#integer) and [decimal](https://www.w3.org/TR/xmlschema-2/#decimal). If the "E" or "e" and the following exponent are omitted, an exponent value of 0 is assumed.

The _special values_ positive and negative infinity and not-a-number have lexical representations `INF`, `-INF` and `NaN`, respectively. Lexical representations for zero may take a positive or negative sign.

For example, `-1E4, 1267.43233E12, 12.78e-2, 12` `, -0, 0` and `INF` are all legal literals for **float**.

##### 3.2.4.2 Canonical representation

The canonical representation for **float** is defined by prohibiting certain options from the [Lexical representation (§3.2.4.1)](https://www.w3.org/TR/xmlschema-2/#float-lexical-representation). Specifically, the exponent must be indicated by "E". Leading zeroes and the preceding optional "+" sign are prohibited in the exponent. If the exponent is zero, it must be indicated by "E0". For the mantissa, the preceding optional "+" sign is prohibited and the decimal point is required. Leading and trailing zeroes are prohibited subject to the following: number representations must be normalized such that there is a single digit which is non-zero to the left of the decimal point and at least a single digit to the right of the decimal point unless the value being represented is zero. The canonical representation for zero is 0.0E0.

##### 3.2.4.3 Constraining facets

**float** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.5 double

[Definition:]  The **double** datatype is patterned after the IEEE double-precision 64-bit floating point type [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754). The basic [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **double** consists of the values _m × 2^e_, where _m_ is an integer whose absolute value is less than _2^53_, and _e_ is an integer between -1075 and 970, inclusive. In addition to the basic [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) described above, the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **double** also contains the following three _special values_: positive and negative infinity and not-a-number (NaN). The [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) on **double** is: _x < y iff y - x_ is positive for x and y in the value space. Positive infinity is greater than all other non-NaN values. NaN equals itself but is [·incomparable·](https://www.w3.org/TR/xmlschema-2/#dt-incomparable) with (neither greater than nor less than) any other value in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

**Note:**  "Equality" in this Recommendation is defined to be "identity" (i.e., values that are identical in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) are equal and vice versa). Identity must be used for the few operations that are defined in this Recommendation. Applications using any of the datatypes defined in this Recommendation may use different definitions of equality for computational purposes; [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754)-based computation systems are examples. Nothing in this Recommendation should be construed as requiring that such applications use identity as their equality relationship when computing.

Any value [·incomparable·](https://www.w3.org/TR/xmlschema-2/#dt-incomparable) with the value used for the four bounding facets ([·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive), [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive), [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive), and [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive)) will be excluded from the resulting restricted [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space). In particular, when "NaN" is used as a facet value for a bounding facet, since no other **double** values are [·comparable·](https://www.w3.org/TR/xmlschema-2/#dt-comparable) with it, the result is a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) either having NaN as its only member (the inclusive cases) or that is empty (the exclusive cases). If any other value is used for a bounding facet, NaN will be excluded from the resulting restricted [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space); to add NaN back in requires union with the NaN-only space.

This datatype differs from that of [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754) in that there is only one NaN and only one zero. This makes the equality and ordering of values in the data space differ from that of [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754) only in that for schema purposes NaN = NaN.

A literal in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) representing a decimal number _d_ maps to the normalized value in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **double** that is closest to _d_; if _d_ is exactly halfway between two such values then the even value is chosen. This is the _best approximation_ of _d_ ([[Clinger, WD (1990)]](https://www.w3.org/TR/xmlschema-2/#clinger1990), [[Gay, DM (1990)]](https://www.w3.org/TR/xmlschema-2/#gay1990)), which is more accurate than the mapping required by [[IEEE 754-1985]](https://www.w3.org/TR/xmlschema-2/#ieee754).

##### 3.2.5.1 Lexical representation

**double** values have a lexical representation consisting of a mantissa followed, optionally, by the character "E" or "e", followed by an exponent. The exponent [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be an integer. The mantissa must be a decimal number. The representations for exponent and mantissa must follow the lexical rules for [integer](https://www.w3.org/TR/xmlschema-2/#integer) and [decimal](https://www.w3.org/TR/xmlschema-2/#decimal). If the "E" or "e" and the following exponent are omitted, an exponent value of 0 is assumed.

The _special values_ positive and negative infinity and not-a-number have lexical representations `INF`, `-INF` and `NaN`, respectively. Lexical representations for zero may take a positive or negative sign.

For example, `-1E4, 1267.43233E12, 12.78e-2, 12` `, -0, 0` and `INF` are all legal literals for **double**.

##### 3.2.5.2 Canonical representation

The canonical representation for **double** is defined by prohibiting certain options from the [Lexical representation (§3.2.5.1)](https://www.w3.org/TR/xmlschema-2/#double-lexical-representation). Specifically, the exponent must be indicated by "E". Leading zeroes and the preceding optional "+" sign are prohibited in the exponent. If the exponent is zero, it must be indicated by "E0". For the mantissa, the preceding optional "+" sign is prohibited and the decimal point is required. Leading and trailing zeroes are prohibited subject to the following: number representations must be normalized such that there is a single digit which is non-zero to the left of the decimal point and at least a single digit to the right of the decimal point unless the value being represented is zero. The canonical representation for zero is 0.0E0.

##### 3.2.5.3 Constraining facets

**double** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.6 duration

[Definition:]   **duration** represents a duration of time. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **duration** is a six-dimensional space where the coordinates designate the Gregorian year, month, day, hour, minute, and second components defined in § 5.5.3.2 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601), respectively. These components are ordered in their significance by their order of appearance i.e. as year, month, day, hour, minute, and second.

**Note:**

All [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) support year values with a minimum of 4 digits (i.e., `YYYY`) and a minimum fractional second precision of milliseconds or three decimal digits (i.e. `s.sss`). However, [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) set an application-defined limit on the maximum number of digits they are prepared to support in these two cases, in which case that application-defined maximum number [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be clearly documented.

##### 3.2.6.1 Lexical representation

The lexical representation for **duration** is the [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601) extended format P_n_Y_n_ M_n_DT_n_H _n_M_n_S, where _n_Y represents the number of years, _n_M the number of months, _n_D the number of days, 'T' is the date/time separator, _n_H the number of hours, _n_M the number of minutes and _n_S the number of seconds. The number of seconds can include decimal digits to arbitrary precision.

The values of the Year, Month, Day, Hour and Minutes components are not restricted but allow an arbitrary unsigned integer, i.e., an integer that conforms to the pattern `[0-9]+`.. Similarly, the value of the Seconds component allows an arbitrary unsigned decimal. Following [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601), at least one digit must follow the decimal point if it appears. That is, the value of the Seconds component must conform to the pattern `[0-9]+(\.[0-9]+)?`. Thus, the lexical representation of **duration** does not follow the alternative format of § 5.5.3.2.1 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601).

An optional preceding minus sign ('-') is allowed, to indicate a negative duration. If the sign is omitted a positive duration is indicated. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

For example, to indicate a duration of 1 year, 2 months, 3 days, 10 hours, and 30 minutes, one would write: `P1Y2M3DT10H30M`. One could also indicate a duration of minus 120 days as: `-P120D`.

Reduced precision and truncated representations of this format are allowed provided they conform to the following:

- If the number of years, months, days, hours, minutes, or seconds in any expression equals zero, the number and its corresponding designator [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) be omitted. However, at least one number and its designator [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be present.
- The seconds part [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) have a decimal fraction.
- The designator 'T' must be absent if and only if all of the time items are absent. The designator 'P' must always be present.

For example, P1347Y, P1347M and P1Y2MT2H are all allowed; P0Y1347M and P0Y1347M0D are allowed. P-1347M is not allowed although -P1347M is allowed. P1Y2MT is not allowed.

##### 3.2.6.2 Order relation on duration

In general, the [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) on **duration** is a partial order since there is no determinate relationship between certain durations such as one month (P1M) and 30 days (P30D). The [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) of two **duration** values _x_ and _y_ is _x < y iff s+x < s+y_ for each qualified [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) _s_ in the list below. These values for _s_ cause the greatest deviations in the addition of dateTimes and durations. Addition of durations to time instants is defined in [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes).

- 1696-09-01T00:00:00Z
- 1697-02-01T00:00:00Z
- 1903-03-01T00:00:00Z
- 1903-07-01T00:00:00Z

The following table shows the strongest relationship that can be determined between example durations. The symbol <> means that the order relation is indeterminate. Note that because of leap-seconds, a seconds field can vary from 59 to 60. However, because of the way that addition is defined in [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes), they are still totally ordered.

||Relation|   |   |   |   |   |   |
|---|---|---|---|---|---|---|---|
|P**1Y**|> P**364D**|<> P**365D**||   |   |<> P**366D**|< P**367D**|
|P**1M**|> P**27D**|<> P**28D**|<> P**29D**|   |<> P**30D**|<> P**31D**|< P**32D**|
|P**5M**|> P**149D**|<> P**150D**|<> P**151D**|<> P**152D**|   |<> P**153D**|< P**154D**|

Implementations are free to optimize the computation of the ordering relationship. For example, the following table can be used to compare durations of a small number of months against days.

||Months|1|2|3|4|5|6|7|8|9|10|11|12|13|...|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
|Days|Minimum|28|59|89|120|150|181|212|242|273|303|334|365|393|...|
|Maximum|31|62|92|123|153|184|215|245|276|306|337|366|397|...|

##### 3.2.6.3 Facet Comparison for durations

In comparing **duration** values with [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) and [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) facet values indeterminate comparisons should be considered as "false".

##### 3.2.6.4 Totally ordered durations

Certain derived datatypes of durations can be guaranteed have a total order. For this, they must have fields from only one row in the list below and the time zone must either be required or prohibited.

- year, month
- day, hour, minute, second

For example, a datatype could be defined to correspond to the [[SQL]](https://www.w3.org/TR/xmlschema-2/#SQL) datatype Year-Month interval that required a four digit year field and a two digit month field but required all other fields to be unspecified. This datatype could be defined as below and would have a total order.

<simpleType name='SQL-Year-Month-Interval'>
    <restriction base='duration'>
      <pattern value='P\p{Nd}{4}Y\p{Nd}{2}M'/>
    </restriction>
</simpleType>

##### 3.2.6.5 Constraining facets

**duration** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.7 dateTime

[Definition:]   **dateTime** values may be viewed as objects with integer-valued year, month, day, hour and minute properties, a decimal-valued second property, and a boolean timezoned property. Each such object also has one decimal-valued method or computed property, timeOnTimeline, whose value is always a decimal number; the values are dimensioned in seconds, the integer 0 is 0001-01-01T00:00:00 and the value of timeOnTimeline for other **dateTime** values is computed using the Gregorian algorithm as modified for leap-seconds. The timeOnTimeline values form two related "timelines", one for timezoned values and one for non-timezoned values. Each timeline is a copy of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [decimal](https://www.w3.org/TR/xmlschema-2/#decimal), with integers given units of seconds.

The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **dateTime** is closely related to the dates and times described in ISO 8601. For clarity, the text above specifies a particular origin point for the timeline. It should be noted, however, that schema processors need not expose the timeOnTimeline value to schema users, and there is no requirement that a timeline-based implementation use the particular origin described here in its internal representation. Other interpretations of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) which lead to the same results (i.e., are isomorphic) are of course acceptable.

All timezoned times are Coordinated Universal Time (UTC, sometimes called "Greenwich Mean Time"). Other timezones indicated in lexical representations are converted to UTC during conversion of literals to values. "Local" or untimezoned times are presumed to be the time in the timezone of some unspecified locality as prescribed by the appropriate legal authority; currently there are no legally prescribed timezones which are durations whose magnitude is greater than 14 hours. The value of each numeric-valued property (other than timeOnTimeline) is limited to the maximum value within the interval determined by the next-higher property. For example, the day value can never be 32, and cannot even be 29 for month 02 and year 2002 (February 2002).

**Note:**

The date and time datatypes described in this recommendation were inspired by [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). '0001' is the lexical representation of the year 1 of the Common Era (1 CE, sometimes written "AD 1" or "1 AD"). There is no year 0, and '0000' is not a valid lexical representation. '-0001' is the lexical representation of the year 1 Before Common Era (1 BCE, sometimes written "1 BC").

Those using this (1.0) version of this Recommendation to represent negative years should be aware that the interpretation of lexical representations beginning with a `'-'` is likely to change in subsequent versions.

[[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601) makes no mention of the year 0; in [[ISO 8601:1998 Draft Revision]](https://www.w3.org/TR/xmlschema-2/#ISO8601-1998) the form '0000' was disallowed and this recommendation disallows it as well. However, [[ISO 8601:2000 Second Edition]](https://www.w3.org/TR/xmlschema-2/#ISO8601-2000), which became available just as we were completing version 1.0, allows the form '0000', representing the year 1 BCE. A number of external commentators have also suggested that '0000' be allowed, as the lexical representation for 1 BCE, which is the normal usage in astronomical contexts. It is the intention of the XML Schema Working Group to allow '0000' as a lexical representation in the **dateTime**, **date**, **gYear**, and **gYearMonth** datatypes in a subsequent version of this Recommendation. '0000' will be the lexical representation of 1 BCE (which is a leap year), '-0001' will become the lexical representation of 2 BCE (not 1 BCE as in this (1.0) version), '-0002' of 3 BCE, etc.

**Note:** See the conformance note in [(§3.2.6)](https://www.w3.org/TR/xmlschema-2/#year-sec-conformance) which applies to this datatype as well.

##### 3.2.7.1 Lexical representation

The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **dateTime** consists of finite-length sequences of characters of the form: `'-'? yyyy '-' mm '-' dd 'T' hh ':' mm ':' ss ('.' s+)? (zzzzzz)?`, where

- '-'? _yyyy_ is a four-or-more digit optionally negative-signed numeral that represents the year; if more than four digits, leading zeros are prohibited, and '0000' is prohibited (see the Note above [(§3.2.7)](https://www.w3.org/TR/xmlschema-2/#year-zero); also note that a plus sign is **not** permitted);
- the remaining '-'s are separators between parts of the date portion;
- the first _mm_ is a two-digit numeral that represents the month;
- _dd_ is a two-digit numeral that represents the day;
- 'T' is a separator indicating that time-of-day follows;
- _hh_ is a two-digit numeral that represents the hour; '24' is permitted if the minutes and seconds represented are zero, and the **dateTime** value so represented is the first instant of the following day (the hour property of a **dateTime** object in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) cannot have a value greater than 23);
- ':' is a separator between parts of the time-of-day portion;
- the second _mm_ is a two-digit numeral that represents the minute;
- _ss_ is a two-integer-digit numeral that represents the whole seconds;
- '.' _s+_ (if present) represents the fractional seconds;
- _zzzzzz_ (if present) represents the timezone (as described below).

For example, 2002-10-10T12:00:00-05:00 (noon on 10 October 2002, Central Daylight Savings Time as well as Eastern Standard Time in the U.S.) is 2002-10-10T17:00:00Z, five hours later than 2002-10-10T12:00:00Z.

For further guidance on arithmetic with **dateTime**s and durations, see [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes).

##### 3.2.7.2 Canonical representation

Except for trailing fractional zero digits in the seconds representation, '24:00:00' time representations, and timezone (for timezoned values), the mapping from literals to values is one-to-one. Where there is more than one possible representation, the canonical representation is as follows:

- The 2-digit numeral representing the hour must not be '`24`';
- The fractional second string, if present, must not end in '`0`';
- for timezoned values, the timezone must be represented with '`Z`' (All timezoned **dateTime** values are UTC.).

##### 3.2.7.3 Timezones

Timezones are durations with (integer-valued) hour and minute properties (with the hour magnitude limited to at most 14, and the minute magnitude limited to at most 59, except that if the hour magnitude is 14, the minute value must be 0); they may be both positive or both negative.

The lexical representation of a timezone is a string of the form: `(('+' | '-') hh ':' mm) | 'Z'`, where

- _hh_ is a two-digit numeral (with leading zeros as required) that represents the hours,
- _mm_ is a two-digit numeral that represents the minutes,
- '+' indicates a nonnegative duration,
- '-' indicates a nonpositive duration.

The mapping so defined is one-to-one, except that '+00:00', '-00:00', and 'Z' all represent the same zero-length duration timezone, UTC; 'Z' is its canonical representation.

When a timezone is added to a UTC **dateTime**, the result is the date and time "in that timezone". For example, 2002-10-10T12:00:00+05:00 is 2002-10-10T07:00:00Z and 2002-10-10T00:00:00+05:00 is 2002-10-09T19:00:00Z.

##### 3.2.7.4 Order relation on dateTime

**dateTime** value objects on either timeline are totally ordered by their timeOnTimeline values; between the two timelines, **dateTime** value objects are ordered by their timeOnTimeline values when their timeOnTimeline values differ by more than fourteen hours, with those whose difference is a duration of 14 hours or less being [·incomparable·](https://www.w3.org/TR/xmlschema-2/#dt-incomparable).

In general, the [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) on **dateTime** is a partial order since there is no determinate relationship between certain instants. For example, there is no determinate ordering between (a) 2000-01-20T12:00:00 and (b) 2000-01-20T12:00:00**Z**. Based on timezones currently in use, (c) could vary from 2000-01-20T12:00:00+12:00 to 2000-01-20T12:00:00-13:00. It is, however, possible for this range to expand or contract in the future, based on local laws. Because of this, the following definition uses a somewhat broader range of indeterminate values: +14:00..-14:00.

The following definition uses the notation S[year] to represent the year field of S, S[month] to represent the month field, and so on. The notation (Q & "-14:00") means adding the timezone -14:00 to Q, where Q did not already have a timezone. _This is a logical explanation of the process. Actual implementations are free to optimize as long as they produce the same results._

The ordering between two **dateTime**s P and Q is defined by the following algorithm:

A.Normalize P and Q. That is, if there is a timezone present, but it is not Z, convert it to Z using the addition operation defined in [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes)

- Thus 2000-03-04T23:00:00+03:00 normalizes to 2000-03-04T20:00:00Z

B. If P and Q either both have a time zone or both do not have a time zone, compare P and Q field by field from the year field down to the second field, and return a result as soon as it can be determined. That is:

1. For each i in {year, month, day, hour, minute, second}
    1. If P[i] and Q[i] are both not specified, continue to the next i
    2. If P[i] is not specified and Q[i] is, or vice versa, stop and return P <> Q
    3. If P[i] < Q[i], stop and return P < Q
    4. If P[i] > Q[i], stop and return P > Q
2. Stop and return P = Q

C.Otherwise, if P contains a time zone and Q does not, compare as follows:

1. P < Q if P < (Q with time zone +14:00)
2. P > Q if P > (Q with time zone -14:00)
3. P <> Q otherwise, that is, if (Q with time zone +14:00) < P < (Q with time zone -14:00)

D. Otherwise, if P does not contain a time zone and Q does, compare as follows:

1. P < Q if (P with time zone -14:00) < Q.
2. P > Q if (P with time zone +14:00) > Q.
3. P <> Q otherwise, that is, if (P with time zone +14:00) < Q < (P with time zone -14:00)

Examples:

|Determinate|Indeterminate|
|:-:|:-:|
|2000-01-15T00:00:00 **<** 2000-02-15T00:00:00|2000-01-01T12:00:00 **<>** 1999-12-31T23:00:00Z|
|2000-01-15T12:00:00 **<** 2000-01-16T12:00:00Z|2000-01-16T12:00:00 **<>** 2000-01-16T12:00:00Z|
||2000-01-16T00:00:00 **<>** 2000-01-16T12:00:00Z|

##### 3.2.7.5 Totally ordered dateTimes

Certain derived types from **dateTime** can be guaranteed have a total order. To do so, they must require that a specific set of fields are always specified, and that remaining fields (if any) are always unspecified. For example, the date datatype without time zone is defined to contain exactly year, month, and day. Thus dates without time zone have a total order among themselves.

##### 3.2.7.6 Constraining facets

**dateTime** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.8 time

[Definition:]  **time** represents an instant of time that recurs every day. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **time** is the space of _time of day_ values as defined in § 5.3 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Specifically, it is a set of zero-duration daily time instances.

Since the lexical representation allows an optional time zone indicator, **time** values are partially ordered because it may not be able to determine the order of two values one of which has a time zone and the other does not. The order relation on **time** values is the [Order relation on dateTime (§3.2.7.4)](https://www.w3.org/TR/xmlschema-2/#dateTime-order) using an arbitrary date. See also [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes). Pairs of **time** values with or without time zone indicators are totally ordered.

**Note:** See the conformance note in [(§3.2.6)](https://www.w3.org/TR/xmlschema-2/#year-sec-conformance) which applies to the seconds part of this datatype as well.

##### 3.2.8.1 Lexical representation

The lexical representation for **time** is the left truncated lexical representation for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime): hh:mm:ss.sss with optional following time zone indicator. For example, to indicate 1:20 pm for Eastern Standard Time which is 5 hours behind Coordinated Universal Time (UTC), one would write: 13:20:00-05:00. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

##### 3.2.8.2 Canonical representation

The canonical representation for **time** is defined by prohibiting certain options from the [Lexical representation (§3.2.8.1)](https://www.w3.org/TR/xmlschema-2/#time-lexical-repr). Specifically, either the time zone must be omitted or, if present, the time zone must be Coordinated Universal Time (UTC) indicated by a "Z". Additionally, the canonical representation for midnight is 00:00:00.

##### 3.2.8.3 Constraining facets

**time** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.9 date

[Definition:]   The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **date** consists of top-open intervals of exactly one day in length on the timelines of [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), beginning on the beginning moment of each day (in each timezone), i.e. '00:00:00', up to but not including '24:00:00' (which is identical with '00:00:00' of the next day). For nontimezoned values, the top-open intervals disjointly cover the nontimezoned timeline, one per day. For timezoned values, the intervals begin at every minute and therefore overlap.

A "date object" is an object with year, month, and day properties just like those of [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) objects, plus an optional _timezone-valued_ timezone property. (As with values of [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) timezones are a special case of durations.) Just as a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) object corresponds to a point on one of the timelines, a **date** object corresponds to an interval on one of the two timelines as just described.

Timezoned **date** values track the starting moment of their day, as determined by their timezone; said timezone is generally recoverable for canonical representations. [Definition:]   The **recoverable timezone** is that duration which is the result of subtracting the first moment (or any moment) of the timezoned **date** from the first moment (or the corresponding moment) UTC on the same **date**. [·recoverable timezone·](https://www.w3.org/TR/xmlschema-2/#recoverable-timezone)s are always durations between '+12:00' and '-11:59'. This "timezone normalization" (which follows automatically from the definition of the **date** [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)) is explained more in [Lexical representation (§3.2.9.1)](https://www.w3.org/TR/xmlschema-2/#date-lexical-representation).

For example: the first moment of 2002-10-10+13:00 is 2002-10-10T00:00:00+13, which is 2002-10-09T11:00:00Z, which is also the first moment of 2002-10-09-11:00. Therefore 2002-10-10+13:00 is 2002-10-09-11:00; _they are the same interval_.

**Note:**  For most timezones, either the first moment or last moment of the day (a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) value, always UTC) will have a **date** portion different from that of the **date** itself! However, noon of that **date** (the midpoint of the interval) in that (normalized) timezone will always have the same **date** portion as the **date** itself, even when that noon point in time is normalized to UTC. For example, 2002-10-10-05:00 begins during 2002-10-09Z and 2002-10-10+05:00 ends during 2002-10-11Z, but noon of both 2002-10-10-05:00 and 2002-10-10+05:00 falls in the interval which is 2002-10-10Z.

**Note:** See the conformance note in [(§3.2.6)](https://www.w3.org/TR/xmlschema-2/#year-sec-conformance) which applies to the year part of this datatype as well.

##### 3.2.9.1 Lexical representation

For the following discussion, let the "date portion" of a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) or **date** object be an object similar to a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) or **date** object, with similar year, month, and day properties, but no others, having the same value for these properties as the original [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) or **date** object.

The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **date** consists of finite-length sequences of characters of the form: `'-'? yyyy '-' mm '-' dd zzzzzz?` where the **date** and optional timezone are represented exactly the same way as they are for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime). The first moment of the interval is that represented by: `'-' yyyy '-' mm '-' dd 'T00:00:00' zzzzzz?` and the least upper bound of the interval is the timeline point represented (noncanonically) by: `'-' yyyy '-' mm '-' dd 'T24:00:00' zzzzzz?`.

**Note:**  The [·recoverable timezone·](https://www.w3.org/TR/xmlschema-2/#recoverable-timezone) of a **date** will always be a duration between '+12:00' and '11:59'. Timezone lexical representations, as explained for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), can range from '+14:00' to '-14:00'. The result is that literals of **date**s with very large or very negative timezones will map to a "normalized" **date** value with a [·recoverable timezone·](https://www.w3.org/TR/xmlschema-2/#recoverable-timezone) different from that represented in the original representation, and a matching difference of +/- 1 day in the **date** itself.

##### 3.2.9.2 Canonical representation

Given a member of the **date** [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), the **date** portion of the canonical representation (the entire representation for nontimezoned values, and all but the timezone representation for timezoned values) is always the **date** portion of the [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) canonical representation of the interval midpoint (the [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) representation, truncated on the right to eliminate 'T' and all following characters). For timezoned values, append the canonical representation of the [·recoverable timezone·](https://www.w3.org/TR/xmlschema-2/#recoverable-timezone).

#### 3.2.10 gYearMonth

[Definition:]   **gYearMonth** represents a specific gregorian month in a specific gregorian year. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **gYearMonth** is the set of Gregorian calendar months as defined in § 5.2.1 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Specifically, it is a set of one-month long, non-periodic instances e.g. 1999-10 to represent the whole month of 1999-10, independent of how many days this month has.

Since the lexical representation allows an optional time zone indicator, **gYearMonth** values are partially ordered because it may not be possible to unequivocally determine the order of two values one of which has a time zone and the other does not. If **gYearMonth** values are considered as periods of time, the order relation on **gYearMonth** values is the order relation on their starting instants. This is discussed in [Order relation on dateTime (§3.2.7.4)](https://www.w3.org/TR/xmlschema-2/#dateTime-order). See also [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes). Pairs of **gYearMonth** values with or without time zone indicators are totally ordered.

**Note:**  Because month/year combinations in one calendar only rarely correspond to month/year combinations in other calendars, values of this type are not, in general, convertible to simple values corresponding to month/year combinations in other calendars. This type should therefore be used with caution in contexts where conversion to other calendars is desired.

**Note:** See the conformance note in [(§3.2.6)](https://www.w3.org/TR/xmlschema-2/#year-sec-conformance) which applies to the year part of this datatype as well.

##### 3.2.10.1 Lexical representation

The lexical representation for **gYearMonth** is the reduced (right truncated) lexical representation for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime): CCYY-MM. No left truncation is allowed. An optional following time zone qualifier is allowed. To accommodate year values outside the range from 0001 to 9999, additional digits can be added to the left of this representation and a preceding "-" sign is allowed.

For example, to indicate the month of May 1999, one would write: 1999-05. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

##### 3.2.10.2 Constraining facets

**gYearMonth** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.11 gYear

[Definition:]   **gYear** represents a gregorian calendar year. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **gYear** is the set of Gregorian calendar years as defined in § 5.2.1 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Specifically, it is a set of one-year long, non-periodic instances e.g. lexical 1999 to represent the whole year 1999, independent of how many months and days this year has.

Since the lexical representation allows an optional time zone indicator, **gYear** values are partially ordered because it may not be possible to unequivocally determine the order of two values one of which has a time zone and the other does not. If **gYear** values are considered as periods of time, the order relation on **gYear** values is the order relation on their starting instants. This is discussed in [Order relation on dateTime (§3.2.7.4)](https://www.w3.org/TR/xmlschema-2/#dateTime-order). See also [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes). Pairs of **gYear** values with or without time zone indicators are totally ordered.

**Note:**  Because years in one calendar only rarely correspond to years in other calendars, values of this type are not, in general, convertible to simple values corresponding to years in other calendars. This type should therefore be used with caution in contexts where conversion to other calendars is desired.

**Note:** See the conformance note in [(§3.2.6)](https://www.w3.org/TR/xmlschema-2/#year-sec-conformance) which applies to the year part of this datatype as well.

##### 3.2.11.1 Lexical representation

The lexical representation for **gYear** is the reduced (right truncated) lexical representation for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime): CCYY. No left truncation is allowed. An optional following time zone qualifier is allowed as for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime). To accommodate year values outside the range from 0001 to 9999, additional digits can be added to the left of this representation and a preceding "-" sign is allowed.

For example, to indicate 1999, one would write: 1999. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

##### 3.2.11.2 Constraining facets

**gYear** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.12 gMonthDay

[Definition:]   **gMonthDay** is a gregorian date that recurs, specifically a day of the year such as the third of May. Arbitrary recurring dates are not supported by this datatype. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **gMonthDay** is the set of _calendar dates_, as defined in § 3 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Specifically, it is a set of one-day long, annually periodic instances.

Since the lexical representation allows an optional time zone indicator, **gMonthDay** values are partially ordered because it may not be possible to unequivocally determine the order of two values one of which has a time zone and the other does not. If **gMonthDay** values are considered as periods of time, in an arbitrary leap year, the order relation on **gMonthDay** values is the order relation on their starting instants. This is discussed in [Order relation on dateTime (§3.2.7.4)](https://www.w3.org/TR/xmlschema-2/#dateTime-order). See also [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes). Pairs of **gMonthDay** values with or without time zone indicators are totally ordered.

**Note:**  Because day/month combinations in one calendar only rarely correspond to day/month combinations in other calendars, values of this type do not, in general, have any straightforward or intuitive representation in terms of most other calendars. This type should therefore be used with caution in contexts where conversion to other calendars is desired.

##### 3.2.12.1 Lexical representation

The lexical representation for **gMonthDay** is the left truncated lexical representation for [date](https://www.w3.org/TR/xmlschema-2/#date): --MM-DD. An optional following time zone qualifier is allowed as for [date](https://www.w3.org/TR/xmlschema-2/#date). No preceding sign is allowed. No other formats are allowed. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

This datatype can be used to represent a specific day in a month. To say, for example, that my birthday occurs on the 14th of September ever year.

##### 3.2.12.2 Constraining facets

**gMonthDay** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.13 gDay

[Definition:]   **gDay** is a gregorian day that recurs, specifically a day of the month such as the 5th of the month. Arbitrary recurring days are not supported by this datatype. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **gDay** is the space of a set of _calendar dates_ as defined in § 3 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Specifically, it is a set of one-day long, monthly periodic instances.

This datatype can be used to represent a specific day of the month. To say, for example, that I get my paycheck on the 15th of each month.

Since the lexical representation allows an optional time zone indicator, **gDay** values are partially ordered because it may not be possible to unequivocally determine the order of two values one of which has a time zone and the other does not. If **gDay** values are considered as periods of time, in an arbitrary month that has 31 days, the order relation on **gDay** values is the order relation on their starting instants. This is discussed in [Order relation on dateTime (§3.2.7.4)](https://www.w3.org/TR/xmlschema-2/#dateTime-order). See also [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes). Pairs of **gDay** values with or without time zone indicators are totally ordered.

**Note:**  Because days in one calendar only rarely correspond to days in other calendars, values of this type do not, in general, have any straightforward or intuitive representation in terms of most other calendars. This type should therefore be used with caution in contexts where conversion to other calendars is desired.

##### 3.2.13.1 Lexical representation

The lexical representation for **gDay** is the left truncated lexical representation for [date](https://www.w3.org/TR/xmlschema-2/#date): ---DD . An optional following time zone qualifier is allowed as for [date](https://www.w3.org/TR/xmlschema-2/#date). No preceding sign is allowed. No other formats are allowed. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

##### 3.2.13.2 Constraining facets

**gDay** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.14 gMonth

[Definition:]   **gMonth** is a gregorian month that recurs every year. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **gMonth** is the space of a set of _calendar months_ as defined in § 3 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Specifically, it is a set of one-month long, yearly periodic instances.

This datatype can be used to represent a specific month. To say, for example, that Thanksgiving falls in the month of November.

Since the lexical representation allows an optional time zone indicator, **gMonth** values are partially ordered because it may not be possible to unequivocally determine the order of two values one of which has a time zone and the other does not. If **gMonth** values are considered as periods of time, the order relation on **gMonth** is the order relation on their starting instants. This is discussed in [Order relation on dateTime (§3.2.7.4)](https://www.w3.org/TR/xmlschema-2/#dateTime-order). See also [Adding durations to dateTimes (§E)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-dateTimes). Pairs of **gMonth** values with or without time zone indicators are totally ordered.

**Note:**  Because months in one calendar only rarely correspond to months in other calendars, values of this type do not, in general, have any straightforward or intuitive representation in terms of most other calendars. This type should therefore be used with caution in contexts where conversion to other calendars is desired.

##### 3.2.14.1 Lexical representation

The lexical representation for **gMonth** is the left and right truncated lexical representation for [date](https://www.w3.org/TR/xmlschema-2/#date): --MM. An optional following time zone qualifier is allowed as for [date](https://www.w3.org/TR/xmlschema-2/#date). No preceding sign is allowed. No other formats are allowed. See also [ISO 8601 Date and Time Formats (§D)](https://www.w3.org/TR/xmlschema-2/#isoformats).

##### 3.2.14.2 Constraining facets

**gMonth** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.2.15 hexBinary

[Definition:]   **hexBinary** represents arbitrary hex-encoded binary data. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **hexBinary** is the set of finite-length sequences of binary octets.

##### 3.2.15.1 Lexical Representation

**hexBinary** has a lexical representation where each binary octet is encoded as a character tuple, consisting of two hexadecimal digits ([0-9a-fA-F]) representing the octet code. For example, "0FB7" is a _hex_ encoding for the 16-bit integer 4023 (whose binary representation is 111110110111).

##### 3.2.15.2 Canonical Representation

The canonical representation for **hexBinary** is defined by prohibiting certain options from the [Lexical Representation (§3.2.15.1)](https://www.w3.org/TR/xmlschema-2/#hexBinary-lexical-representation). Specifically, the lower case hexadecimal digits ([a-f]) are not allowed.

##### 3.2.15.3 Constraining facets

**hexBinary** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

#### 3.2.16 base64Binary

[Definition:]   **base64Binary** represents Base64-encoded arbitrary binary data. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **base64Binary** is the set of finite-length sequences of binary octets. For **base64Binary** data the entire binary stream is encoded using the Base64 Alphabet in [[RFC 2045]](https://www.w3.org/TR/xmlschema-2/#RFC2045).

The lexical forms of **base64Binary** values are limited to the 65 characters of the Base64 Alphabet defined in [[RFC 2045]](https://www.w3.org/TR/xmlschema-2/#RFC2045), i.e., `a-z`, `A-Z`, `0-9`, the plus sign (+), the forward slash (/) and the equal sign (=), together with the characters defined in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML) as white space. No other characters are allowed.

For compatibility with older mail gateways, [[RFC 2045]](https://www.w3.org/TR/xmlschema-2/#RFC2045) suggests that base64 data should have lines limited to at most 76 characters in length. This line-length limitation is not mandated in the lexical forms of **base64Binary** data and must not be enforced by XML Schema processors.

The lexical space of **base64Binary** is given by the following grammar (the notation is that used in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML)); legal lexical forms must match the **Base64Binary** production.

`Base64Binary  ::=  ((B64S B64S B64S B64S)*                        ((B64S B64S B64S B64) |                         (B64S B64S B16S '=') |                         (B64S B04S '=' #x20? '=')))?      B64S         ::= B64 #x20?      B16S         ::= B16 #x20?      B04S         ::= B04 #x20?``      B04         ::=  [AQgw]   B16         ::=  [AEIMQUYcgkosw048]   B64         ::=  [A-Za-z0-9+/]`

Note that this grammar requires the number of non-whitespace characters in the lexical form to be a multiple of four, and for equals signs to appear only at the end of the lexical form; strings which do not meet these constraints are not legal lexical forms of **base64Binary** because they cannot successfully be decoded by base64 decoders.

**Note:** The above definition of the lexical space is more restrictive than that given in [[RFC 2045]](https://www.w3.org/TR/xmlschema-2/#RFC2045) as regards whitespace -- this is not an issue in practice. Any string compatible with the RFC can occur in an element or attribute validated by this type, because the [·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace) facet of this type is fixed to _collapse_, which means that all leading and trailing whitespace will be stripped, and all internal whitespace collapsed to single space characters, _before_ the above grammar is enforced.

The canonical lexical form of a **base64Binary** data value is the base64 encoding of the value which matches the Canonical-base64Binary production in the following grammar:

`Canonical-base64Binary  ::=  (B64 B64 B64 B64)*                                  ((B64 B64 B16 '=') | (B64 B04 '=='))?`

**Note:** For some values the canonical form defined above does not conform to [[RFC 2045]](https://www.w3.org/TR/xmlschema-2/#RFC2045), which requires breaking with linefeeds at appropriate intervals.

The length of a **base64Binary** value is the number of octets it contains. This may be calculated from the lexical form by removing whitespace and padding characters and performing the calculation shown in the pseudo-code below:

`lex2    := killwhitespace(lexform)    -- remove whitespace characters   lex3    := strip_equals(lex2)         -- strip padding characters at end   length  := floor (length(lex3) * 3 / 4)         -- calculate length`

Note on encoding: [[RFC 2045]](https://www.w3.org/TR/xmlschema-2/#RFC2045) explicitly references US-ASCII encoding. However, decoding of **base64Binary** data in an XML entity is to be performed on the Unicode characters obtained after character encoding processing as specified by [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML)

##### 3.2.16.1 Constraining facets

**base64Binary** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

#### 3.2.17 anyURI

[Definition:]   **anyURI** represents a Uniform Resource Identifier Reference (URI). An **anyURI** value can be absolute or relative, and may have an optional fragment identifier (i.e., it may be a URI Reference). This type should be used to specify the intention that the value fulfills the role of a URI as defined by [[RFC 2396]](https://www.w3.org/TR/xmlschema-2/#RFC2396), as amended by [[RFC 2732]](https://www.w3.org/TR/xmlschema-2/#RFC2732).

The mapping from **anyURI** values to URIs is as defined by the URI reference escaping procedure defined in Section 5.4 [Locator Attribute](https://www.w3.org/TR/2001/REC-xlink-20010627/#link-locators) of [[XML Linking Language]](https://www.w3.org/TR/xmlschema-2/#XLink) (see also Section 8 [Character Encoding in URI References](https://www.w3.org/TR/2001/WD-charmod-20010126/#sec-URIs) of [[Character Model]](https://www.w3.org/TR/xmlschema-2/#CharMod)). This means that a wide range of internationalized resource identifiers can be specified when an **anyURI** is called for, and still be understood as URIs per [[RFC 2396]](https://www.w3.org/TR/xmlschema-2/#RFC2396), as amended by [[RFC 2732]](https://www.w3.org/TR/xmlschema-2/#RFC2732), where appropriate to identify resources.

**Note:**  Section 5.4 [Locator Attribute](https://www.w3.org/TR/2001/REC-xlink-20010627/#link-locators) of [[XML Linking Language]](https://www.w3.org/TR/xmlschema-2/#XLink) requires that relative URI references be absolutized as defined in [[XML Base]](https://www.w3.org/TR/xmlschema-2/#XBase) before use. This is an XLink-specific requirement and is not appropriate for XML Schema, since neither the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) nor the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI) type are restricted to absolute URIs. Accordingly absolutization must not be performed by schema processors as part of schema validation.

**Note:**  Each URI scheme imposes specialized syntax rules for URIs in that scheme, including restrictions on the syntax of allowed fragment identifiers. Because it is impractical for processors to check that a value is a context-appropriate URI reference, this specification follows the lead of [[RFC 2396]](https://www.w3.org/TR/xmlschema-2/#RFC2396) (as amended by [[RFC 2732]](https://www.w3.org/TR/xmlschema-2/#RFC2732)) in this matter: such rules and restrictions are not part of type validity and are not checked by [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors. Thus in practice the above definition imposes only very modest obligations on [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors.

##### 3.2.17.1 Lexical representation

The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **anyURI** is finite-length character sequences which, when the algorithm defined in Section 5.4 of [[XML Linking Language]](https://www.w3.org/TR/xmlschema-2/#XLink) is applied to them, result in strings which are legal URIs according to [[RFC 2396]](https://www.w3.org/TR/xmlschema-2/#RFC2396), as amended by [[RFC 2732]](https://www.w3.org/TR/xmlschema-2/#RFC2732).

**Note:**  Spaces are, in principle, allowed in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **anyURI**, however, their use is highly discouraged (unless they are encoded by %20).

##### 3.2.17.2 Constraining facets

**anyURI** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

#### 3.2.18 QName

[Definition:]   **QName** represents [XML qualified names](https://www.w3.org/TR/1999/REC-xml-names-19990114/#dt-qname). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **QName** is the set of tuples {[namespace name](https://www.w3.org/TR/1999/REC-xml-names-19990114/#dt-NSName), [local part](https://www.w3.org/TR/1999/REC-xml-names-19990114/#dt-localname)}, where [namespace name](https://www.w3.org/TR/1999/REC-xml-names-19990114/#dt-NSName) is an [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI) and [local part](https://www.w3.org/TR/1999/REC-xml-names-19990114/#dt-localname) is an [NCName](https://www.w3.org/TR/xmlschema-2/#NCName). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **QName** is the set of strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [QName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-QName) production of [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS).

**Note:**  The mapping between literals in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) and values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **QName** requires a namespace declaration to be in scope for the context in which **QName** is used.

##### 3.2.18.1 Constraining facets

**QName** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

The use of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length), [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength) and [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) on datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [QName](https://www.w3.org/TR/xmlschema-2/#QName) is deprecated. Future versions of this specification may remove these facets for this datatype.

#### 3.2.19 NOTATION

[Definition:]   **NOTATION** represents the [NOTATION](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-NotationType) attribute type from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **NOTATION** is the set of [QName](https://www.w3.org/TR/xmlschema-2/#QName)s of notations declared in the current schema. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **NOTATION** is the set of all names of [notations](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#declare-notation) declared in the current schema (in the form of [QName](https://www.w3.org/TR/xmlschema-2/#QName)s).

**Schema Component Constraint: enumeration facet value required for NOTATION**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for **NOTATION** to be used directly in a schema. Only datatypes that are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **NOTATION** by specifying a value for [·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration) can be used in a schema.

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **NOTATION** should be used only on attributes and should only be used in schemas with no target namespace.

##### 3.2.19.1 Constraining facets

**NOTATION** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

The use of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length), [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength) and [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) on datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION) is deprecated. Future versions of this specification may remove these facets for this datatype.

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#built-in-primitive-datatypes)3.3 Derived datatypes

        3.3.1 [normalizedString](https://www.w3.org/TR/xmlschema-2/#normalizedString)
        3.3.2 [token](https://www.w3.org/TR/xmlschema-2/#token)
        3.3.3 [language](https://www.w3.org/TR/xmlschema-2/#language)
        3.3.4 [NMTOKEN](https://www.w3.org/TR/xmlschema-2/#NMTOKEN)
        3.3.5 [NMTOKENS](https://www.w3.org/TR/xmlschema-2/#NMTOKENS)
        3.3.6 [Name](https://www.w3.org/TR/xmlschema-2/#Name)
        3.3.7 [NCName](https://www.w3.org/TR/xmlschema-2/#NCName)
        3.3.8 [ID](https://www.w3.org/TR/xmlschema-2/#ID)
        3.3.9 [IDREF](https://www.w3.org/TR/xmlschema-2/#IDREF)
        3.3.10 [IDREFS](https://www.w3.org/TR/xmlschema-2/#IDREFS)
        3.3.11 [ENTITY](https://www.w3.org/TR/xmlschema-2/#ENTITY)
        3.3.12 [ENTITIES](https://www.w3.org/TR/xmlschema-2/#ENTITIES)
        3.3.13 [integer](https://www.w3.org/TR/xmlschema-2/#integer)
        3.3.14 [nonPositiveInteger](https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger)
        3.3.15 [negativeInteger](https://www.w3.org/TR/xmlschema-2/#negativeInteger)
        3.3.16 [long](https://www.w3.org/TR/xmlschema-2/#long)
        3.3.17 [int](https://www.w3.org/TR/xmlschema-2/#int)
        3.3.18 [short](https://www.w3.org/TR/xmlschema-2/#short)
        3.3.19 [byte](https://www.w3.org/TR/xmlschema-2/#byte)
        3.3.20 [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)
        3.3.21 [unsignedLong](https://www.w3.org/TR/xmlschema-2/#unsignedLong)
        3.3.22 [unsignedInt](https://www.w3.org/TR/xmlschema-2/#unsignedInt)
        3.3.23 [unsignedShort](https://www.w3.org/TR/xmlschema-2/#unsignedShort)
        3.3.24 [unsignedByte](https://www.w3.org/TR/xmlschema-2/#unsignedByte)
        3.3.25 [positiveInteger](https://www.w3.org/TR/xmlschema-2/#positiveInteger)

This section gives conceptual definitions for all [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes defined by this specification. The XML representation used to define [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes (whether [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) or [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived)) is given in section [XML Representation of Simple Type Definition Schema Components (§4.1.2)](https://www.w3.org/TR/xmlschema-2/#xr-defn) and the complete definitions of the [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in)  [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes are provided in Appendix A [Schema for Datatype Definitions (normative) (§A)](https://www.w3.org/TR/xmlschema-2/#schema).

#### 3.3.1 normalizedString

[Definition:]   **normalizedString** represents white space normalized strings. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **normalizedString** is the set of strings that do not contain the carriage return (#xD), line feed (#xA) nor tab (#x9) characters. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **normalizedString** is the set of strings that do not contain the carriage return (#xD), line feed (#xA) nor tab (#x9) characters. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **normalizedString** is [string](https://www.w3.org/TR/xmlschema-2/#string).

##### 3.3.1.1 Constraining facets

**normalizedString** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.1.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **normalizedString**:

- [token](https://www.w3.org/TR/xmlschema-2/#token)

#### 3.3.2 token

[Definition:]   **token** represents tokenized strings. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **token** is the set of strings that do not contain the carriage return (#xD), line feed (#xA) nor tab (#x9) characters, that have no leading or trailing spaces (#x20) and that have no internal sequences of two or more spaces. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **token** is the set of strings that do not contain the carriage return (#xD), line feed (#xA) nor tab (#x9) characters, that have no leading or trailing spaces (#x20) and that have no internal sequences of two or more spaces. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **token** is [normalizedString](https://www.w3.org/TR/xmlschema-2/#normalizedString).

##### 3.3.2.1 Constraining facets

**token** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.2.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **token**:

- [language](https://www.w3.org/TR/xmlschema-2/#language)
- [NMTOKEN](https://www.w3.org/TR/xmlschema-2/#NMTOKEN)
- [Name](https://www.w3.org/TR/xmlschema-2/#Name)

#### 3.3.3 language

[Definition:]   **language** represents natural language identifiers as defined by by [[RFC 3066]](https://www.w3.org/TR/xmlschema-2/#RFC3066) . The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **language** is the set of all strings that are valid language identifiers as defined [[RFC 3066]](https://www.w3.org/TR/xmlschema-2/#RFC3066) . The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **language** is the set of all strings that conform to the pattern `[a-zA-Z]{1,8}(-[a-zA-Z0-9]{1,8})*` . The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **language** is [token](https://www.w3.org/TR/xmlschema-2/#token).

##### 3.3.3.1 Constraining facets

**language** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

#### 3.3.4 NMTOKEN

[Definition:]   **NMTOKEN** represents the [NMTOKEN attribute type](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **NMTOKEN** is the set of tokens that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [Nmtoken](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Nmtoken) production in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **NMTOKEN** is the set of strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [Nmtoken](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Nmtoken) production in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **NMTOKEN** is [token](https://www.w3.org/TR/xmlschema-2/#token).

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **NMTOKEN** should be used only on attributes.

##### 3.3.4.1 Constraining facets

**NMTOKEN** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.4.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **NMTOKEN**:

- [NMTOKENS](https://www.w3.org/TR/xmlschema-2/#NMTOKENS)

#### 3.3.5 NMTOKENS

[Definition:]   **NMTOKENS** represents the [NMTOKENS attribute type](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **NMTOKENS** is the set of finite, non-zero-length sequences of [·NMTOKEN·](https://www.w3.org/TR/xmlschema-2/#dt-NMTOKEN)s. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **NMTOKENS** is the set of space-separated lists of tokens, of which each token is in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of [NMTOKEN](https://www.w3.org/TR/xmlschema-2/#NMTOKEN). The [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) of **NMTOKENS** is [NMTOKEN](https://www.w3.org/TR/xmlschema-2/#NMTOKEN).

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **NMTOKENS** should be used only on attributes.

##### 3.3.5.1 Constraining facets

**NMTOKENS** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)

#### 3.3.6 Name

[Definition:]   **Name** represents [XML Names](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-name). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **Name** is the set of all strings which [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [Name](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Name) production of [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **Name** is the set of all strings which [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [Name](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Name) production of [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **Name** is [token](https://www.w3.org/TR/xmlschema-2/#token).

##### 3.3.6.1 Constraining facets

**Name** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.6.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **Name**:

- [NCName](https://www.w3.org/TR/xmlschema-2/#NCName)

#### 3.3.7 NCName

[Definition:]   **NCName** represents XML "non-colonized" Names. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **NCName** is the set of all strings which [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production of [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **NCName** is the set of all strings which [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production of [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **NCName** is [Name](https://www.w3.org/TR/xmlschema-2/#Name).

##### 3.3.7.1 Constraining facets

**NCName** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.7.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **NCName**:

- [ID](https://www.w3.org/TR/xmlschema-2/#ID)
- [IDREF](https://www.w3.org/TR/xmlschema-2/#IDREF)
- [ENTITY](https://www.w3.org/TR/xmlschema-2/#ENTITY)

#### 3.3.8 ID

[Definition:]   **ID** represents the [ID attribute type](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **ID** is the set of all strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **ID** is the set of all strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **ID** is [NCName](https://www.w3.org/TR/xmlschema-2/#NCName).

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **ID** should be used only on attributes.

##### 3.3.8.1 Constraining facets

**ID** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

#### 3.3.9 IDREF

[Definition:]   **IDREF** represents the [IDREF attribute type](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **IDREF** is the set of all strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **IDREF** is the set of strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **IDREF** is [NCName](https://www.w3.org/TR/xmlschema-2/#NCName).

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) this datatype should be used only on attributes.

##### 3.3.9.1 Constraining facets

**IDREF** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.9.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **IDREF**:

- [IDREFS](https://www.w3.org/TR/xmlschema-2/#IDREFS)

#### 3.3.10 IDREFS

[Definition:]   **IDREFS** represents the [IDREFS attribute type](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **IDREFS** is the set of finite, non-zero-length sequences of [IDREF](https://www.w3.org/TR/xmlschema-2/#IDREF)s. The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **IDREFS** is the set of space-separated lists of tokens, of which each token is in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of [IDREF](https://www.w3.org/TR/xmlschema-2/#IDREF). The [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) of **IDREFS** is [IDREF](https://www.w3.org/TR/xmlschema-2/#IDREF).

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **IDREFS** should be used only on attributes.

##### 3.3.10.1 Constraining facets

**IDREFS** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)

#### 3.3.11 ENTITY

[Definition:]   **ENTITY** represents the [ENTITY](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) attribute type from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **ENTITY** is the set of all strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS) and have been declared as an [unparsed entity](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-unparsed) in a [document type definition](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-doctype). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **ENTITY** is the set of all strings that [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) the [NCName](https://www.w3.org/TR/1999/REC-xml-names-19990114/#NT-NCName) production in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS). The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **ENTITY** is [NCName](https://www.w3.org/TR/xmlschema-2/#NCName).

**Note:**  The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **ENTITY** is scoped to a specific instance document.

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **ENTITY** should be used only on attributes.

##### 3.3.11.1 Constraining facets

**ENTITY** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)

##### 3.3.11.2 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **ENTITY**:

- [ENTITIES](https://www.w3.org/TR/xmlschema-2/#ENTITIES)

#### 3.3.12 ENTITIES

[Definition:]   **ENTITIES** represents the [ENTITIES attribute type](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-TokenizedType) from [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **ENTITIES** is the set of finite, non-zero-length sequences of [·ENTITY·](https://www.w3.org/TR/xmlschema-2/#dt-ENTITY)s that have been declared as [unparsed entities](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-unparsed) in a [document type definition](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-doctype). The [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of **ENTITIES** is the set of space-separated lists of tokens, of which each token is in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of [ENTITY](https://www.w3.org/TR/xmlschema-2/#ENTITY). The [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) of **ENTITIES** is [ENTITY](https://www.w3.org/TR/xmlschema-2/#ENTITY).

**Note:**  The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **ENTITIES** is scoped to a specific instance document.

For compatibility (see [Terminology (§1.4)](https://www.w3.org/TR/xmlschema-2/#terminology)) **ENTITIES** should be used only on attributes.

##### 3.3.12.1 Constraining facets

**ENTITIES** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
- [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
- [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)

#### 3.3.13 integer

[Definition:]   **integer** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [decimal](https://www.w3.org/TR/xmlschema-2/#decimal) by fixing the value of [·fractionDigits·](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits) to be 0and disallowing the trailing decimal point. This results in the standard mathematical concept of the integer numbers. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **integer** is the infinite set {...,-2,-1,0,1,2,...}. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **integer** is [decimal](https://www.w3.org/TR/xmlschema-2/#decimal).

##### 3.3.13.1 Lexical representation

**integer** has a lexical representation consisting of a finite-length sequence of decimal digits (#x30-#x39) with an optional leading sign. If the sign is omitted, "+" is assumed. For example: -1, 0, 12678967543233, +100000.

##### 3.3.13.2 Canonical representation

The canonical representation for **integer** is defined by prohibiting certain options from the [Lexical representation (§3.3.13.1)](https://www.w3.org/TR/xmlschema-2/#integer-lexical-representation). Specifically, the preceding optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.13.3 Constraining facets

**integer** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.13.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **integer**:

- [nonPositiveInteger](https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger)
- [long](https://www.w3.org/TR/xmlschema-2/#long)
- [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)

#### 3.3.14 nonPositiveInteger

[Definition:]   **nonPositiveInteger** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [integer](https://www.w3.org/TR/xmlschema-2/#integer) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 0. This results in the standard mathematical concept of the non-positive integers. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **nonPositiveInteger** is the infinite set {...,-2,-1,0}. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **nonPositiveInteger** is [integer](https://www.w3.org/TR/xmlschema-2/#integer).

##### 3.3.14.1 Lexical representation

**nonPositiveInteger** has a lexical representation consisting of an optional preceding sign followed by a finite-length sequence of decimal digits (#x30-#x39). The sign may be "+" or may be omitted only for lexical forms denoting zero; in all other lexical forms, the negative sign ("-") must be present. For example: -1, 0, -12678967543233, -100000.

##### 3.3.14.2 Canonical representation

The canonical representation for **nonPositiveInteger** is defined by prohibiting certain options from the [Lexical representation (§3.3.14.1)](https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger-lexical-representation). In the canonical form for zero, the sign must be omitted. Leading zeroes are prohibited.

##### 3.3.14.3 Constraining facets

**nonPositiveInteger** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.14.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **nonPositiveInteger**:

- [negativeInteger](https://www.w3.org/TR/xmlschema-2/#negativeInteger)

#### 3.3.15 negativeInteger

[Definition:]   **negativeInteger** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [nonPositiveInteger](https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be -1. This results in the standard mathematical concept of the negative integers. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **negativeInteger** is the infinite set {...,-2,-1}. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **negativeInteger** is [nonPositiveInteger](https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger).

##### 3.3.15.1 Lexical representation

**negativeInteger** has a lexical representation consisting of a negative sign ("-") followed by a finite-length sequence of decimal digits (#x30-#x39). For example: -1, -12678967543233, -100000.

##### 3.3.15.2 Canonical representation

The canonical representation for **negativeInteger** is defined by prohibiting certain options from the [Lexical representation (§3.3.15.1)](https://www.w3.org/TR/xmlschema-2/#negativeInteger-lexical-representation). Specifically, leading zeroes are prohibited.

##### 3.3.15.3 Constraining facets

**negativeInteger** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.3.16 long

[Definition:]   **long** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [integer](https://www.w3.org/TR/xmlschema-2/#integer) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 9223372036854775807 and [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be -9223372036854775808. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **long** is [integer](https://www.w3.org/TR/xmlschema-2/#integer).

##### 3.3.16.1 Lexical representation

**long** has a lexical representation consisting of an optional sign followed by a finite-length sequence of decimal digits (#x30-#x39). If the sign is omitted, "+" is assumed. For example: -1, 0, 12678967543233, +100000.

##### 3.3.16.2 Canonical representation

The canonical representation for **long** is defined by prohibiting certain options from the [Lexical representation (§3.3.16.1)](https://www.w3.org/TR/xmlschema-2/#long-lexical-representation). Specifically, the the optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.16.3 Constraining facets

**long** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.16.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **long**:

- [int](https://www.w3.org/TR/xmlschema-2/#int)

#### 3.3.17 int

[Definition:]   **int** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [long](https://www.w3.org/TR/xmlschema-2/#long) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 2147483647 and [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be -2147483648. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **int** is [long](https://www.w3.org/TR/xmlschema-2/#long).

##### 3.3.17.1 Lexical representation

**int** has a lexical representation consisting of an optional sign followed by a finite-length sequence of decimal digits (#x30-#x39). If the sign is omitted, "+" is assumed. For example: -1, 0, 126789675, +100000.

##### 3.3.17.2 Canonical representation

The canonical representation for **int** is defined by prohibiting certain options from the [Lexical representation (§3.3.17.1)](https://www.w3.org/TR/xmlschema-2/#int-lexical-representation). Specifically, the the optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.17.3 Constraining facets

**int** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.17.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **int**:

- [short](https://www.w3.org/TR/xmlschema-2/#short)

#### 3.3.18 short

[Definition:]   **short** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [int](https://www.w3.org/TR/xmlschema-2/#int) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 32767 and [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be -32768. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **short** is [int](https://www.w3.org/TR/xmlschema-2/#int).

##### 3.3.18.1 Lexical representation

**short** has a lexical representation consisting of an optional sign followed by a finite-length sequence of decimal digits (#x30-#x39). If the sign is omitted, "+" is assumed. For example: -1, 0, 12678, +10000.

##### 3.3.18.2 Canonical representation

The canonical representation for **short** is defined by prohibiting certain options from the [Lexical representation (§3.3.18.1)](https://www.w3.org/TR/xmlschema-2/#short-lexical-representation). Specifically, the the optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.18.3 Constraining facets

**short** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.18.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **short**:

- [byte](https://www.w3.org/TR/xmlschema-2/#byte)

#### 3.3.19 byte

[Definition:]   **byte** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [short](https://www.w3.org/TR/xmlschema-2/#short) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 127 and [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be -128. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **byte** is [short](https://www.w3.org/TR/xmlschema-2/#short).

##### 3.3.19.1 Lexical representation

**byte** has a lexical representation consisting of an optional sign followed by a finite-length sequence of decimal digits (#x30-#x39). If the sign is omitted, "+" is assumed. For example: -1, 0, 126, +100.

##### 3.3.19.2 Canonical representation

The canonical representation for **byte** is defined by prohibiting certain options from the [Lexical representation (§3.3.19.1)](https://www.w3.org/TR/xmlschema-2/#byte-lexical-representation). Specifically, the the optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.19.3 Constraining facets

**byte** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.3.20 nonNegativeInteger

[Definition:]   **nonNegativeInteger** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [integer](https://www.w3.org/TR/xmlschema-2/#integer) by setting the value of [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be 0. This results in the standard mathematical concept of the non-negative integers. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **nonNegativeInteger** is the infinite set {0,1,2,...}. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **nonNegativeInteger** is [integer](https://www.w3.org/TR/xmlschema-2/#integer).

##### 3.3.20.1 Lexical representation

**nonNegativeInteger** has a lexical representation consisting of an optional sign followed by a finite-length sequence of decimal digits (#x30-#x39). If the sign is omitted, the positive sign ("+") is assumed. If the sign is present, it must be "+" except for lexical forms denoting zero, which may be preceded by a positive ("+") or a negative ("-") sign. For example: 1, 0, 12678967543233, +100000.

##### 3.3.20.2 Canonical representation

The canonical representation for **nonNegativeInteger** is defined by prohibiting certain options from the [Lexical representation (§3.3.20.1)](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger-lexical-representation). Specifically, the the optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.20.3 Constraining facets

**nonNegativeInteger** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.20.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **nonNegativeInteger**:

- [unsignedLong](https://www.w3.org/TR/xmlschema-2/#unsignedLong)
- [positiveInteger](https://www.w3.org/TR/xmlschema-2/#positiveInteger)

#### 3.3.21 unsignedLong

[Definition:]   **unsignedLong** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 18446744073709551615. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **unsignedLong** is [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

##### 3.3.21.1 Lexical representation

**unsignedLong** has a lexical representation consisting of a finite-length sequence of decimal digits (#x30-#x39). For example: 0, 12678967543233, 100000.

##### 3.3.21.2 Canonical representation

The canonical representation for **unsignedLong** is defined by prohibiting certain options from the [Lexical representation (§3.3.21.1)](https://www.w3.org/TR/xmlschema-2/#unsignedLong-lexical-representation). Specifically, leading zeroes are prohibited.

##### 3.3.21.3 Constraining facets

**unsignedLong** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.21.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **unsignedLong**:

- [unsignedInt](https://www.w3.org/TR/xmlschema-2/#unsignedInt)

#### 3.3.22 unsignedInt

[Definition:]   **unsignedInt** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [unsignedLong](https://www.w3.org/TR/xmlschema-2/#unsignedLong) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 4294967295. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **unsignedInt** is [unsignedLong](https://www.w3.org/TR/xmlschema-2/#unsignedLong).

##### 3.3.22.1 Lexical representation

**unsignedInt** has a lexical representation consisting of a finite-length sequence of decimal digits (#x30-#x39). For example: 0, 1267896754, 100000.

##### 3.3.22.2 Canonical representation

The canonical representation for **unsignedInt** is defined by prohibiting certain options from the [Lexical representation (§3.3.22.1)](https://www.w3.org/TR/xmlschema-2/#unsignedInt-lexical-representation). Specifically, leading zeroes are prohibited.

##### 3.3.22.3 Constraining facets

**unsignedInt** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.22.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **unsignedInt**:

- [unsignedShort](https://www.w3.org/TR/xmlschema-2/#unsignedShort)

#### 3.3.23 unsignedShort

[Definition:]   **unsignedShort** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [unsignedInt](https://www.w3.org/TR/xmlschema-2/#unsignedInt) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 65535. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **unsignedShort** is [unsignedInt](https://www.w3.org/TR/xmlschema-2/#unsignedInt).

##### 3.3.23.1 Lexical representation

**unsignedShort** has a lexical representation consisting of a finite-length sequence of decimal digits (#x30-#x39). For example: 0, 12678, 10000.

##### 3.3.23.2 Canonical representation

The canonical representation for **unsignedShort** is defined by prohibiting certain options from the [Lexical representation (§3.3.23.1)](https://www.w3.org/TR/xmlschema-2/#unsignedShort-lexical-representation). Specifically, the leading zeroes are prohibited.

##### 3.3.23.3 Constraining facets

**unsignedShort** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

##### 3.3.23.4 Derived datatypes

The following [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatypes are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **unsignedShort**:

- [unsignedByte](https://www.w3.org/TR/xmlschema-2/#unsignedByte)

#### 3.3.24 unsignedByte

[Definition:]   **unsignedByte** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [unsignedShort](https://www.w3.org/TR/xmlschema-2/#unsignedShort) by setting the value of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) to be 255. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **unsignedByte** is [unsignedShort](https://www.w3.org/TR/xmlschema-2/#unsignedShort).

##### 3.3.24.1 Lexical representation

**unsignedByte** has a lexical representation consisting of a finite-length sequence of decimal digits (#x30-#x39). For example: 0, 126, 100.

##### 3.3.24.2 Canonical representation

The canonical representation for **unsignedByte** is defined by prohibiting certain options from the [Lexical representation (§3.3.24.1)](https://www.w3.org/TR/xmlschema-2/#unsignedByte-lexical-representation). Specifically, leading zeroes are prohibited.

##### 3.3.24.3 Constraining facets

**unsignedByte** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

#### 3.3.25 positiveInteger

[Definition:]   **positiveInteger** is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger) by setting the value of [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be 1. This results in the standard mathematical concept of the positive integer numbers. The [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of **positiveInteger** is the infinite set {1,2,...}. The [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) of **positiveInteger** is [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

##### 3.3.25.1 Lexical representation

**positiveInteger** has a lexical representation consisting of an optional positive sign ("+") followed by a finite-length sequence of decimal digits (#x30-#x39). For example: 1, 12678967543233, +100000.

##### 3.3.25.2 Canonical representation

The canonical representation for **positiveInteger** is defined by prohibiting certain options from the [Lexical representation (§3.3.25.1)](https://www.w3.org/TR/xmlschema-2/#positiveInteger-lexical-representation). Specifically, the optional "+" sign is prohibited and leading zeroes are prohibited.

##### 3.3.25.3 Constraining facets

**positiveInteger** has the following [·constraining facets·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet):

- [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
- [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)
- [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
- [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
- [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
- [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
- [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
- [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
- [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)

## 4 Datatype components

The following sections provide full details on the properties and significance of each kind of schema component involved in datatype definitions. For each property, the kinds of values it is allowed to have is specified. Any property not identified as optional is required to be present; optional properties which are not present have [absent](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-null) as their value. Any property identified as a having a set, subset or [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) value may have an empty value unless this is explicitly ruled out: this is not the same as [absent](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-null). Any property value identified as a superset or a subset of some set may be equal to that set, unless a proper superset or subset is explicitly called for.

For more information on the notion of datatype (schema) components, see [Schema Component Details](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#components) of [[XML Schema Part 1: Structures]](https://www.w3.org/TR/xmlschema-2/#structural-schemas).

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#rf-fund-facets)4.1 Simple Type Definition

        4.1.1 [The Simple Type Definition Schema Component](https://www.w3.org/TR/xmlschema-2/#dc-defn)
        4.1.2 [XML Representation of Simple Type Definition Schema Components](https://www.w3.org/TR/xmlschema-2/#xr-defn)
        4.1.3 [Constraints on XML Representation of Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#defn-rep-constr)
        4.1.4 [Simple Type Definition Validation Rules](https://www.w3.org/TR/xmlschema-2/#defn-validation-rules)
        4.1.5 [Constraints on Simple Type Definition Schema Components](https://www.w3.org/TR/xmlschema-2/#defn-coss)
        4.1.6 [Simple Type Definition for anySimpleType](https://www.w3.org/TR/xmlschema-2/#anySimpleType-component)

Simple Type definitions provide for:

- Establishing the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of a datatype, through the combined set of [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s specified in the definition;
- Attaching a unique name (actually a [QName](https://www.w3.org/TR/xmlschema-2/#QName)) to the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space).

#### 4.1.1 The Simple Type Definition Schema Component

The Simple Type Definition schema component has the following properties:

Schema Component: [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#datatype)

{name}

Optional. An NCName as defined by [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS).

{target namespace}

Either [absent](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-null) or a namespace name, as defined in [[Namespaces in XML]](https://www.w3.org/TR/xmlschema-2/#XMLNS).

{variety}

One of {_atomic_, _list_, _union_}. Depending on the value of [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety), further properties are defined as follows:

atomic

{primitive type definition}

A [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatype definition).

list

{item type definition}

An [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) simple type definition.

union

{member type definitions}

A non-empty sequence of simple type definitions.

{facets}

A possibly empty set of [Facets (§2.4)](https://www.w3.org/TR/xmlschema-2/#facets).

{fundamental facets}

A set of [Fundamental facets (§2.4.1)](https://www.w3.org/TR/xmlschema-2/#fundamental-facets)

{base type definition}

If the datatype has been [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction) then the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component from which it is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived), otherwise the [Simple Type Definition for anySimpleType (§4.1.6)](https://www.w3.org/TR/xmlschema-2/#anySimpleType-component).

{final}

A subset of _{restriction, list, union}_.

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

Datatypes are identified by their [{name}](https://www.w3.org/TR/xmlschema-2/#defn-name) and [{target namespace}](https://www.w3.org/TR/xmlschema-2/#defn-target-namespace). Except for anonymous datatypes (those with no [{name}](https://www.w3.org/TR/xmlschema-2/#defn-name)), datatype definitions [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be uniquely identified within a schema.

If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) then the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the datatype defined will be a subset of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) (which is a subset of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive)). If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) then the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the datatype defined will be the set of finite-length sequence of values from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{item type definition}](https://www.w3.org/TR/xmlschema-2/#defn-itemType). If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) then the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the datatype defined will be the union of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of each datatype in [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes).

If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) then the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) must be [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic). If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) then the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) of [{item type definition}](https://www.w3.org/TR/xmlschema-2/#defn-itemType) must be either [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union). If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) then [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) must be a list of datatype definitions.

The value of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) consists of the set of [·facet·](https://www.w3.org/TR/xmlschema-2/#dt-facet)s specified directly in the datatype definition unioned with the possibly empty set of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

The value of [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) consists of the set of [·fundamental facet·](https://www.w3.org/TR/xmlschema-2/#dt-fundamental-facet)s and their values.

If [{final}](https://www.w3.org/TR/xmlschema-2/#defn-final) is the empty set then the type can be used in deriving other types; the explicit values _restriction_, _list_ and _union_ prevent further derivations by [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction), [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) and [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) respectively.

#### 4.1.2 XML Representation of Simple Type Definition Schema Components

The XML representation for a [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) schema component is a [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `simpleType` Element Information Item

<simpleType
  final = (#all | List of (list | union | restriction))
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **name** = [NCName](https://www.w3.org/TR/xmlschema-2/#NCName)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?, ([restriction](https://www.w3.org/TR/xmlschema-2/#element-restriction) | [list](https://www.w3.org/TR/xmlschema-2/#element-list) | [union](https://www.w3.org/TR/xmlschema-2/#element-union)))
</simpleType>

|[Datatype Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{name}](https://www.w3.org/TR/xmlschema-2/#defn-name)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `name` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise [null](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-null)\|<br>\|[{final}](https://www.w3.org/TR/xmlschema-2/#defn-final)\|A set corresponding to the [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `final` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise the [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `finalDefault` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) of the ancestor [schema](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-schema) element information item, if present, otherwise the empty string, as follows:<br><br>the empty string<br><br>the empty set;<br><br>`#all`<br><br>_{restriction, list, union}_;<br><br>_otherwise_<br><br>a set with members drawn from the set above, each being present or absent depending on whether the string contains an equivalently named space-delimited substring.<br><br>**Note:** Although the `finalDefault` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) of [schema](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-schema) may include values other than _restriction_, _list_ or _union_, those values are ignored in the determination of [{final}](https://www.w3.org/TR/xmlschema-2/#defn-final)\|<br>\|[{target namespace}](https://www.w3.org/TR/xmlschema-2/#defn-target-namespace)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `targetNamespace` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) of the parent `schema` element information item.\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotation corresponding to the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information item in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if present, otherwise [null](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-null)\||

A [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatype can be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from a [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatype or another [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatype by one of three means: by _restriction_, by _list_ or by _union_.

##### 4.1.2.1 Derivation by restriction

XML Representation Summary: `restriction` Element Information Item

<restriction
  base = [QName](https://www.w3.org/TR/xmlschema-2/#QName)
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?, ([simpleType](https://www.w3.org/TR/xmlschema-2/#element-simpleType)?, ([minExclusive](https://www.w3.org/TR/xmlschema-2/#element-minExclusive) | [minInclusive](https://www.w3.org/TR/xmlschema-2/#element-minInclusive) | [maxExclusive](https://www.w3.org/TR/xmlschema-2/#element-maxExclusive) | [maxInclusive](https://www.w3.org/TR/xmlschema-2/#element-maxInclusive) | [totalDigits](https://www.w3.org/TR/xmlschema-2/#element-totalDigits) | [fractionDigits](https://www.w3.org/TR/xmlschema-2/#element-fractionDigits) | [length](https://www.w3.org/TR/xmlschema-2/#element-length) | [minLength](https://www.w3.org/TR/xmlschema-2/#element-minLength) | [maxLength](https://www.w3.org/TR/xmlschema-2/#element-maxLength) | [enumeration](https://www.w3.org/TR/xmlschema-2/#element-enumeration) | [whiteSpace](https://www.w3.org/TR/xmlschema-2/#element-whiteSpace) | [pattern](https://www.w3.org/TR/xmlschema-2/#element-pattern))*))
</restriction>

|[Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype)\|<br>\|[{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets)\|The union of the set of [Facets (§2.4)](https://www.w3.org/TR/xmlschema-2/#facets) components resolved to by the facet [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element) merged with [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) from [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype), subject to the Facet Restriction Valid constraints specified in [Facets (§2.4)](https://www.w3.org/TR/xmlschema-2/#facets).\|<br>\|[{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype)\|The [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component resolved to by the [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `base` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) or the [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType) [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), whichever is present.\||

Example

An electronic commerce schema might define a datatype called _Sku_ (the barcode number that appears on products) from the [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatype [string](https://www.w3.org/TR/xmlschema-2/#string) by supplying a value for the [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) facet.

<simpleType name='Sku'>
    <restriction base='string'>
      <pattern value='\d{3}-[A-Z]{2}'/>
    </restriction>
</simpleType>

In this case, _Sku_ is the name of the new [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype, [string](https://www.w3.org/TR/xmlschema-2/#string) is its [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) and [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) is the facet.

##### 4.1.2.2 Derivation by list

XML Representation Summary: `list` Element Information Item

<list
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  itemType = [QName](https://www.w3.org/TR/xmlschema-2/#QName)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?, [simpleType](https://www.w3.org/TR/xmlschema-2/#element-simpleType)?)
</list>

|[Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety)\|list\|<br>\|[{item type definition}](https://www.w3.org/TR/xmlschema-2/#defn-itemType)\|The [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component resolved to by the [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `itemType` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) or the [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType) [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), whichever is present.\||

A [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype must be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype, known as the [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) of the [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype. This yields a datatype whose [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is composed of finite-length sequences of values from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) and whose [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) is composed of space-separated lists of literals of the [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType).

Example

A system might want to store lists of floating point values.

<simpleType name='listOfFloat'>
  <list itemType='float'/>
</simpleType>

In this case, _listOfFloat_ is the name of the new [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype, [float](https://www.w3.org/TR/xmlschema-2/#float) is its [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) and [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) is the derivation method.

As mentioned in [List datatypes (§2.5.1.2)](https://www.w3.org/TR/xmlschema-2/#list-datatypes), when a datatype is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype, the following [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s can be used:

- [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length)
- [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength)
- [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength)
- [·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration)
- [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern)
- [·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace)

regardless of the [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s that are applicable to the [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype that serves as the [·itemType·](https://www.w3.org/TR/xmlschema-2/#dt-itemType) of the [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list).

For each of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length), [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) and [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength), the _unit of length_ is measured in number of list items. The value of [·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace) is fixed to the value _collapse_.

##### 4.1.2.3 Derivation by union

XML Representation Summary: `union` Element Information Item

<union
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  memberTypes = List of [QName](https://www.w3.org/TR/xmlschema-2/#QName)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?, [simpleType](https://www.w3.org/TR/xmlschema-2/#element-simpleType)*)
</union>

|[Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety)\|union\|<br>\|[{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes)\|The sequence of [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) components resolved to by the items in the [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `memberTypes` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if any, in order, followed by the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) components resolved to by the [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType) [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any, in order. If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is _union_ for any [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) components resolved to above, then the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) is replaced by its [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes).\||

A [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype can be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from one or more [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic), [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) or other [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatypes, known as the [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes) of that [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype.

Example

As an example, taken from a typical display oriented text markup language, one might want to express font sizes as an integer between 8 and 72, or with one of the tokens "small", "medium" or "large". The [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) type definition below would accomplish that.

<xsd:attribute name="size">
  <xsd:simpleType>
    <xsd:union>
      <xsd:simpleType>
        <xsd:restriction base="xsd:positiveInteger">
          <xsd:minInclusive value="8"/>
          <xsd:maxInclusive value="72"/>
        </xsd:restriction>
      </xsd:simpleType>
      <xsd:simpleType>
        <xsd:restriction base="xsd:NMTOKEN">
          <xsd:enumeration value="small"/>
          <xsd:enumeration value="medium"/>
          <xsd:enumeration value="large"/>
        </xsd:restriction>
      </xsd:simpleType>
    </xsd:union>
  </xsd:simpleType>
</xsd:attribute>

<p>
<font size='large'>A header</font>
</p>
<p>
<font size='12'>this is a test</font>
</p>

As mentioned in [Union datatypes (§2.5.1.3)](https://www.w3.org/TR/xmlschema-2/#union-datatypes), when a datatype is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype, the only following [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s can be used:

- [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern)
- [·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration)

regardless of the [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s that are applicable to the datatypes that participate in the [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union)

#### 4.1.3 Constraints on XML Representation of Simple Type Definition

**Schema Representation Constraint: Single Facet Value**

Unless otherwise specifically allowed by this specification ([Multiple patterns (§4.3.4.3)](https://www.w3.org/TR/xmlschema-2/#src-multiple-patterns) and [Multiple enumerations (§4.3.5.3)](https://www.w3.org/TR/xmlschema-2/#src-multiple-enumerations)) any given [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet) can only be specifed once within a single derivation step.

**Schema Representation Constraint: itemType attribute or simpleType child**

Either the `itemType` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) or the [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType) [[child]](https://www.w3.org/TR/xml-infoset/#infoitem.element) of the [<list>](https://www.w3.org/TR/xmlschema-2/#element-list) element must be present, but not both.

**Schema Representation Constraint: base attribute or simpleType child**

Either the `base` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) or the `simpleType` [[child]](https://www.w3.org/TR/xml-infoset/#infoitem.element) of the [<restriction>](https://www.w3.org/TR/xmlschema-2/#element-restriction) element must be present, but not both.

**Schema Representation Constraint: memberTypes attribute or simpleType children**

Either the `memberTypes` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute) of the [<union>](https://www.w3.org/TR/xmlschema-2/#element-union) element must be non-empty or there must be at least one `simpleType` [[child]](https://www.w3.org/TR/xml-infoset/#infoitem.element).

#### 4.1.4 Simple Type Definition Validation Rules

**Validation Rule: Facet Valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to a [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet) component if:

1 the value is facet-valid with respect to the particular [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet) as specified below.

**Validation Rule: Datatype Valid**

A string is datatype-valid with respect to a datatype definition if:

1 it [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match)es a literal in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of the datatype, determined as follows:

1.1 if [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) is a member of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets), then the string must be [pattern valid (§4.3.4.4)](https://www.w3.org/TR/xmlschema-2/#cvc-pattern-valid);

1.2 if [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) is not a member of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets), then

1.2.1 if [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) then the string must [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) a literal in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype)

1.2.2 if [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) then the string must be a sequence of space-separated tokens, each of which [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match)es a literal in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of [{item type definition}](https://www.w3.org/TR/xmlschema-2/#defn-itemType)

1.2.3 if [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) then the string must [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) a literal in the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) of at least one member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes)

2 the value denoted by the literal [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match)ed in the previous step is a member of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the datatype, as determined by it being [Facet Valid (§4.1.4)](https://www.w3.org/TR/xmlschema-2/#cvc-facet-valid) with respect to each member of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) (except for [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern)).

#### 4.1.5 Constraints on Simple Type Definition Schema Components

**Schema Component Constraint: applicable facets**

The [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s which are allowed to be members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) are dependent on [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) as specified in the following table:

|[{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype)|applicable [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets)|
|---|---|
|If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [list](https://www.w3.org/TR/xmlschema-2/#dt-list), then|   |
|[all datatypes]|[length](https://www.w3.org/TR/xmlschema-2/#dt-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dt-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dt-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dt-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dt-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace)|
|If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [union](https://www.w3.org/TR/xmlschema-2/#dt-union), then|   |
|[all datatypes]|[pattern](https://www.w3.org/TR/xmlschema-2/#dt-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dt-enumeration)|
|else if [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [atomic](https://www.w3.org/TR/xmlschema-2/#dt-atomic), then|   |
|[string](https://www.w3.org/TR/xmlschema-2/#string)|[length](https://www.w3.org/TR/xmlschema-2/#dc-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|
|[boolean](https://www.w3.org/TR/xmlschema-2/#boolean)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|
|[float](https://www.w3.org/TR/xmlschema-2/#float)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[double](https://www.w3.org/TR/xmlschema-2/#double)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[decimal](https://www.w3.org/TR/xmlschema-2/#decimal)|[totalDigits](https://www.w3.org/TR/xmlschema-2/#dc-totalDigits), [fractionDigits](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[duration](https://www.w3.org/TR/xmlschema-2/#duration)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[time](https://www.w3.org/TR/xmlschema-2/#time)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[date](https://www.w3.org/TR/xmlschema-2/#date)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[gYear](https://www.w3.org/TR/xmlschema-2/#gYear)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[gDay](https://www.w3.org/TR/xmlschema-2/#gDay)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth)|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace), [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive), [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive), [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive), [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)|
|[hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary)|[length](https://www.w3.org/TR/xmlschema-2/#dc-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|
|[base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary)|[length](https://www.w3.org/TR/xmlschema-2/#dc-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|
|[anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI)|[length](https://www.w3.org/TR/xmlschema-2/#dc-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|
|[QName](https://www.w3.org/TR/xmlschema-2/#QName)|[length](https://www.w3.org/TR/xmlschema-2/#dc-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|
|[NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION)|[length](https://www.w3.org/TR/xmlschema-2/#dc-length), [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength), [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength), [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern), [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration), [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace)|

**Schema Component Constraint: list of atomic**

If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), then the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) of [{item type definition}](https://www.w3.org/TR/xmlschema-2/#defn-itemType)  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union).

**Schema Component Constraint: no circular unions**

If [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union), then it is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [{name}](https://www.w3.org/TR/xmlschema-2/#defn-name) and [{target namespace}](https://www.w3.org/TR/xmlschema-2/#defn-target-namespace)  [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match) [{name}](https://www.w3.org/TR/xmlschema-2/#defn-name) and [{target namespace}](https://www.w3.org/TR/xmlschema-2/#defn-target-namespace) of any member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes).

#### 4.1.6 Simple Type Definition for anySimpleType

There is a simple type definition nearly equivalent to the simple version of the [ur-type definition](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-urType) present in every schema by definition. It has the following properties:

Schema Component: [anySimpleType](https://www.w3.org/TR/xmlschema-2/#dt-anySimpleType)

{name}

anySimpleType

{target namespace}

http://www.w3.org/2001/XMLSchema

{basetype definition}

[the ur-type definition](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#ur-type-itself)

{final}

the empty set

{variety}

[absent](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-null)

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#rf-defn) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#rf-facets)4.2 Fundamental Facets

        4.2.1 [equal](https://www.w3.org/TR/xmlschema-2/#equal)
        4.2.2 [ordered](https://www.w3.org/TR/xmlschema-2/#rf-ordered)
        4.2.3 [bounded](https://www.w3.org/TR/xmlschema-2/#rf-bounded)
        4.2.4 [cardinality](https://www.w3.org/TR/xmlschema-2/#rf-cardinality)
        4.2.5 [numeric](https://www.w3.org/TR/xmlschema-2/#rf-numeric)

#### 4.2.1 equal

Every [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) supports the notion of equality, with the following rules:

- for any _a_ and _b_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), either _a_ is equal to _b_, denoted _a = b_, or _a_ is not equal to _b_, denoted _a != b_
- there is no pair _a_ and _b_ from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) such that both _a = b_ and _a != b_
- for all _a_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), _a = a_
- for any _a_ and _b_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), _a = b_ if and only if _b = a_
- for any _a_, _b_ and _c_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), if _a = b_ and _b = c_, then _a = c_
- for any _a_ and _b_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) if _a = b_, then _a_ and _b_ cannot be distinguished (i.e., equality is identity)
- the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of all [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes are disjoint (they do not share any values)

On every datatype, the operation Equal is defined in terms of the equality property of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space): for any values _a, b_ drawn from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), _Equal(a,b)_ is true if _a = b_, and false otherwise.

Note that in consequence of the above:

- given [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _A_ and [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _B_ where _A_ and _B_ are disjoint, every pair of values _a_ from _A_ and _b_ from _B_, _a != b_
- two values which are members of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the same [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatype may always be compared with each other
- if a datatype _T_ is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) from [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes) _A, B, ..._ then the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _T_ is the union of [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of its [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes) _A, B, ..._. Some values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _T_ are also values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _A_. Other values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _T_ will be values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _B_ and so on. Values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _T_ which are also in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _A_ can be compared with other values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _A_ according to the above rules. Similarly for values of type _T_ and _B_ and all the other [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes).
- if a datatype _T'_ is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction) from an atomic datatype _T_ then the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _T'_ is a subset of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of _T_. Values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of _T_ and _T'_ can be compared according to the above rules
- if datatypes _T'_ and _T''_ are [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction) from a common atomic ancestor _T_ then the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of _T'_ and _T''_ may overlap. Values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s of _T'_ and _T''_ can be compared according to the above rules

**Note:**  There is no schema component corresponding to the **equal** [·fundamental facet·](https://www.w3.org/TR/xmlschema-2/#dt-fundamental-facet).

#### 4.2.2 ordered

[Definition:]  An **order relation** on a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is a mathematical relation that imposes a [·total order·](https://www.w3.org/TR/xmlschema-2/#dt-total-order) or a [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order) on the members of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[Definition:]  A [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), and hence a datatype, is said to be **ordered** if there exists an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) defined for that [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[Definition:]   A **partial order** is an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) that is **irreflexive**, **asymmetric** and **transitive**.

A [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order) has the following properties:

- for no _a_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), _a < a_ (irreflexivity)
- for all _a_ and _b_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), _a < b_ implies not(_b < a_) (asymmetry)
- for all _a_, _b_ and _c_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), _a < b_ and _b < c_ implies _a < c_ (transitivity)

The notation _a <> b_ is used to indicate the case when _a != b_ and neither _a < b_ nor _b < a_. For any values _a_ and _b_ from different [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s, _a <> b_.

[Definition:]  When _a <> b_, _a_ and _b_ are **incomparable**,[Definition:]  otherwise they are **comparable**.

[Definition:]   A **total order** is an [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order) such that for no _a_ and _b_ is it the case that _a <> b_.

A [·total order·](https://www.w3.org/TR/xmlschema-2/#dt-total-order) has all of the properties specified above for [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order), plus the following property:

- for all _a_ and _b_ in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), either _a < b_ or _b < a_ or _a = b_

**Note:**  The fact that this specification does not define an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) for some datatype does not mean that some other application cannot treat that datatype as being ordered by imposing its own order relation.

[·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) provides for:

- indicating whether an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) is defined on a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), and if so, whether that [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) is a [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order) or a [·total order·](https://www.w3.org/TR/xmlschema-2/#dt-total-order)

##### 4.2.2.1 The ordered Schema Component

Schema Component: [ordered](https://www.w3.org/TR/xmlschema-2/#dt-ordered)

{value}

One of _{false, partial, total}_.

[{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) depends on [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety), [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) and [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) in the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component in which a [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) component appears as a member of [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets).

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic), [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) is inherited from [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype). For all [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) types [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is as specified in the table in [Fundamental Facets (§C.1)](https://www.w3.org/TR/xmlschema-2/#app-fundamental-facets).

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) is _false_.

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union), [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) is _partial_ unless one of the following:

- If every member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) is derived from a common ancestor other than the simple ur-type, then [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) is the same as that ancestor's **ordered** facet
- If every member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) has a [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) of _false_ for the **ordered** facet, then [{value}](https://www.w3.org/TR/xmlschema-2/#ordered-value) is _false_

#### 4.2.3 bounded

[Definition:]   A value _u_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _U_ is said to be an **inclusive upper bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _U_) if for all _v_ in _V_, _u_ >= _v_.

[Definition:]   A value _u_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _U_ is said to be an **exclusive upper bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _U_) if for all _v_ in _V_, _u_ > _v_.

[Definition:]   A value _l_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _L_ is said to be an **inclusive lower bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _L_) if for all _v_ in _V_, _l_ <= _v_.

[Definition:]   A value _l_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _L_ is said to be an **exclusive lower bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _L_) if for all _v_ in _V_, _l_ < _v_.

[Definition:]  A datatype is **bounded** if its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) has either an [·inclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-upper-bound) or an [·exclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-upper-bound) and either an [·inclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-lower-bound) or an [·exclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-lower-bound).

[·bounded·](https://www.w3.org/TR/xmlschema-2/#dt-bounded) provides for:

- indicating whether a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is [·bounded·](https://www.w3.org/TR/xmlschema-2/#dt-bounded)

##### 4.2.3.1 The bounded Schema Component

Schema Component: [bounded](https://www.w3.org/TR/xmlschema-2/#dt-bounded)

{value}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

[{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) depends on [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety), [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) and [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) in the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component in which a [·bounded·](https://www.w3.org/TR/xmlschema-2/#dt-bounded) component appears as a member of [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets).

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic), if one of [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) or [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) and one of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) or [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) are among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) , then [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _true_; else [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _false_.

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), if [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length) or both of [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength) and [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) are among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets), then [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _true_; else [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _false_.

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union), if [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _true_ for every member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) and all members of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) share a common ancestor, then [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _true_; else [{value}](https://www.w3.org/TR/xmlschema-2/#bounded-value) is _false_.

#### 4.2.4 cardinality

[Definition:]  Every [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) has associated with it the concept of **cardinality**. Some [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s are finite, some are countably infinite while still others could conceivably be uncountably infinite (although no [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) defined by this specification is uncountable infinite). A datatype is said to have the cardinality of its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

It is sometimes useful to categorize [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s (and hence, datatypes) as to their cardinality. There are two significant cases:

- [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s that are finite
- [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s that are countably infinite

[·cardinality·](https://www.w3.org/TR/xmlschema-2/#dt-cardinality) provides for:

- indicating whether the [·cardinality·](https://www.w3.org/TR/xmlschema-2/#dt-cardinality) of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is _finite_ or _countably infinite_

##### 4.2.4.1 The cardinality Schema Component

Schema Component: [cardinality](https://www.w3.org/TR/xmlschema-2/#dt-cardinality)

{value}

One of _{finite, countably infinite}_.

[{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) depends on [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety), [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) and [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) in the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component in which a [·cardinality·](https://www.w3.org/TR/xmlschema-2/#dt-cardinality) component appears as a member of [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets).

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) and [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is _finite_, then [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _finite_.

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) and [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is _countably infinite_ and **either** of the following conditions are true, then [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _finite_; else [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _countably infinite_:

1. one of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length), [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength), [·totalDigits·](https://www.w3.org/TR/xmlschema-2/#dt-totalDigits) is among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets),
2. **all** of the following are true:
    1. one of [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) or [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) is among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets)
    2. one of [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) or [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) is among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets)
    3. **either** of the following are true:
        1. [·fractionDigits·](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits) is among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets)
        2. [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is one of [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gYear](https://www.w3.org/TR/xmlschema-2/#gYear), [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay), [gDay](https://www.w3.org/TR/xmlschema-2/#gDay) or [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth) or any type [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from them

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), if [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length) or both of [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength) and [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) are among [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets), then [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _finite_; else [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _countably infinite_.

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union), if [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _finite_ for every member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes), then [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _finite_; else [{value}](https://www.w3.org/TR/xmlschema-2/#cardinality-value) is _countably infinite_.

#### 4.2.5 numeric

[Definition:]  A datatype is said to be **numeric** if its values are conceptually quantities (in some mathematical number system).

[Definition:]  A datatype whose values are not [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) is said to be **non-numeric**.

[·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) provides for:

- indicating whether a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric)

##### 4.2.5.1 The numeric Schema Component

Schema Component: [numeric](https://www.w3.org/TR/xmlschema-2/#dt-numeric)

{value}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean)

[{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) depends on [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety), [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets), [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes) in the [Simple Type Definition](https://www.w3.org/TR/xmlschema-2/#dc-defn) component in which a [·cardinality·](https://www.w3.org/TR/xmlschema-2/#dt-cardinality) component appears as a member of [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets).

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic), [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is inherited from [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype). For all [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) types [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is as specified in the table in [Fundamental Facets (§C.1)](https://www.w3.org/TR/xmlschema-2/#app-fundamental-facets).

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is _false_.

When [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union), if [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is _true_ for every member of [{member type definitions}](https://www.w3.org/TR/xmlschema-2/#defn-memberTypes), then [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is _true_; else [{value}](https://www.w3.org/TR/xmlschema-2/#numeric-value) is _false_.

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#rf-fund-facets)4.3 Constraining Facets

        4.3.1 [length](https://www.w3.org/TR/xmlschema-2/#rf-length)
        4.3.2 [minLength](https://www.w3.org/TR/xmlschema-2/#rf-minLength)
        4.3.3 [maxLength](https://www.w3.org/TR/xmlschema-2/#rf-maxLength)
        4.3.4 [pattern](https://www.w3.org/TR/xmlschema-2/#rf-pattern)
        4.3.5 [enumeration](https://www.w3.org/TR/xmlschema-2/#rf-enumeration)
        4.3.6 [whiteSpace](https://www.w3.org/TR/xmlschema-2/#rf-whiteSpace)
        4.3.7 [maxInclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxInclusive)
        4.3.8 [maxExclusive](https://www.w3.org/TR/xmlschema-2/#rf-maxExclusive)
        4.3.9 [minExclusive](https://www.w3.org/TR/xmlschema-2/#rf-minExclusive)
        4.3.10 [minInclusive](https://www.w3.org/TR/xmlschema-2/#rf-minInclusive)
        4.3.11 [totalDigits](https://www.w3.org/TR/xmlschema-2/#rf-totalDigits)
        4.3.12 [fractionDigits](https://www.w3.org/TR/xmlschema-2/#rf-fractionDigits)

#### 4.3.1 length

[Definition:]   **length** is the number of _units of length_, where _units of length_ varies depending on the type that is being [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from. The value of **length** [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

For [string](https://www.w3.org/TR/xmlschema-2/#string) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string), **length** is measured in units of [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s as defined in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). For [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI), **length** is measured in units of characters (as for [string](https://www.w3.org/TR/xmlschema-2/#string)). For [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary) and [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from them, **length** is measured in octets (8 bits) of binary data. For datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), **length** is measured in number of list items.

**Note:**  For [string](https://www.w3.org/TR/xmlschema-2/#string) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string), **length** will not always coincide with "string length" as perceived by some users or with the number of storage units in some digital representation. Therefore, care should be taken when specifying a value for **length** and in attempting to infer storage requirements from a given value for **length**.

[·length·](https://www.w3.org/TR/xmlschema-2/#dt-length) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with a specific number of _units of length_, where _units of length_ varies depending on [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype to represent product codes which must be exactly 8 characters in length. By fixing the value of the **length** facet we ensure that types derived from productCode can change or set the values of other facets, such as **pattern**, but cannot change the length.

<simpleType name='productCode'>
   <restriction base='string'>
     <length value='8' fixed='true'/>
   </restriction>
</simpleType>

##### 4.3.1.1 The length Schema Component

Schema Component: [length](https://www.w3.org/TR/xmlschema-2/#dt-length)

{value}

A [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#length-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [length](https://www.w3.org/TR/xmlschema-2/#dc-length) other than [{value}](https://www.w3.org/TR/xmlschema-2/#length-value).

##### 4.3.1.2 XML Representation of length Schema Components

The XML representation for a [length](https://www.w3.org/TR/xmlschema-2/#dc-length) schema component is a [<length>](https://www.w3.org/TR/xmlschema-2/#element-length) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `length` Element Information Item

<length
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</length>

|[length](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#length-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#length-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.1.3 length Validation Rules

**Validation Rule: Length Valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length), determined as follows:

1 if the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) then

1.1 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [string](https://www.w3.org/TR/xmlschema-2/#string) or [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI), then the length of the value, as measured in [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be equal to [{value}](https://www.w3.org/TR/xmlschema-2/#length-value);

1.2 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary) or [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary), then the length of the value, as measured in octets of the binary data, [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be equal to [{value}](https://www.w3.org/TR/xmlschema-2/#length-value);

1.3 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [QName](https://www.w3.org/TR/xmlschema-2/#QName) or [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION), then any [{value}](https://www.w3.org/TR/xmlschema-2/#length-value) is facet-valid.

2 if the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), then the length of the value, as measured in list items, [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be equal to [{value}](https://www.w3.org/TR/xmlschema-2/#length-value)

The use of [·length·](https://www.w3.org/TR/xmlschema-2/#dt-length) on datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [QName](https://www.w3.org/TR/xmlschema-2/#QName) and [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION) is deprecated. Future versions of this specification may remove this facet for these datatypes.

##### 4.3.1.4 Constraints on length Schema Components

**Schema Component Constraint: length and minLength or maxLength**

If [length](https://www.w3.org/TR/xmlschema-2/#dc-length) is a member of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) then

1 It is an error for [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) to be a member of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) unless

1.1 the [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value) of [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) <= the [{value}](https://www.w3.org/TR/xmlschema-2/#length-value) of [length](https://www.w3.org/TR/xmlschema-2/#dc-length) and

1.2 there is type definition from which this one is derived by one or more restriction steps in which [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) has the same [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value) and [length](https://www.w3.org/TR/xmlschema-2/#dc-length) is not specified.

2 It is an error for [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) to be a member of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) unless

2.1 the [{value}](https://www.w3.org/TR/xmlschema-2/#length-value) of [length](https://www.w3.org/TR/xmlschema-2/#dc-length) <= the [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value) of [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) and

2.2 there is type definition from which this one is derived by one or more restriction steps in which [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) has the same [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value) and [length](https://www.w3.org/TR/xmlschema-2/#dc-length) is not specified.

**Schema Component Constraint: length valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [length](https://www.w3.org/TR/xmlschema-2/#dc-length) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#length-value) is not equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#length-value) of the parent [length](https://www.w3.org/TR/xmlschema-2/#dc-length).

#### 4.3.2 minLength

[Definition:]   **minLength** is the minimum number of _units of length_, where _units of length_ varies depending on the type that is being [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from. The value of **minLength**  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

For [string](https://www.w3.org/TR/xmlschema-2/#string) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string), **minLength** is measured in units of [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s as defined in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). For [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary) and [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from them, **minLength** is measured in octets (8 bits) of binary data. For datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), **minLength** is measured in number of list items.

**Note:**  For [string](https://www.w3.org/TR/xmlschema-2/#string) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string), **minLength** will not always coincide with "string length" as perceived by some users or with the number of storage units in some digital representation. Therefore, care should be taken when specifying a value for **minLength** and in attempting to infer storage requirements from a given value for **minLength**.

[·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with at least a specific number of _units of length_, where _units of length_ varies depending on [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which requires strings to have at least one character (i.e., the empty string is not in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of this datatype).

<simpleType name='non-empty-string'>
  <restriction base='string'>
    <minLength value='1'/>
  </restriction>
</simpleType>

##### 4.3.2.1 The minLength Schema Component

Schema Component: [minLength](https://www.w3.org/TR/xmlschema-2/#dt-minLength)

{value}

A [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#minLength-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) other than [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value).

##### 4.3.2.2 XML Representation of minLength Schema Component

The XML representation for a [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) schema component is a [<minLength>](https://www.w3.org/TR/xmlschema-2/#element-minLength) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `minLength` Element Information Item

<minLength
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</minLength>

|[minLength](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#minLength-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.2.3 minLength Validation Rules

**Validation Rule: minLength Valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength), determined as follows:

1 if the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) then

1.1 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [string](https://www.w3.org/TR/xmlschema-2/#string) or [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI), then the length of the value, as measured in [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be greater than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value);

1.2 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary) or [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary), then the length of the value, as measured in octets of the binary data, [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be greater than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value);

1.3 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [QName](https://www.w3.org/TR/xmlschema-2/#QName) or [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION), then any [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value) is facet-valid.

2 if the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), then the length of the value, as measured in list items, [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be greater than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value)

The use of [·minLength·](https://www.w3.org/TR/xmlschema-2/#dt-minLength) on datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [QName](https://www.w3.org/TR/xmlschema-2/#QName) and [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION) is deprecated. Future versions of this specification may remove this facet for these datatypes.

##### 4.3.2.4 Constraints on minLength Schema Components

**Schema Component Constraint: minLength <= maxLength**

If both [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) and [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) are members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets), then the [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value) of [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength)  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be less than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value) of [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength).

**Schema Component Constraint: minLength valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value) is less than the [{value}](https://www.w3.org/TR/xmlschema-2/#minLength-value) of the parent [minLength](https://www.w3.org/TR/xmlschema-2/#dc-minLength).

#### 4.3.3 maxLength

[Definition:]   **maxLength** is the maximum number of _units of length_, where _units of length_ varies depending on the type that is being [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from. The value of **maxLength**  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

For [string](https://www.w3.org/TR/xmlschema-2/#string) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string), **maxLength** is measured in units of [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s as defined in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML). For [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary) and [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from them, **maxLength** is measured in octets (8 bits) of binary data. For datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), **maxLength** is measured in number of list items.

**Note:**  For [string](https://www.w3.org/TR/xmlschema-2/#string) and datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string), **maxLength** will not always coincide with "string length" as perceived by some users or with the number of storage units in some digital representation. Therefore, care should be taken when specifying a value for **maxLength** and in attempting to infer storage requirements from a given value for **maxLength**.

[·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with at most a specific number of _units of length_, where _units of length_ varies depending on [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which might be used to accept form input with an upper limit to the number of characters that are acceptable.

<simpleType name='form-input'>
  <restriction base='string'>
    <maxLength value='50'/>
  </restriction>
</simpleType>

##### 4.3.3.1 The maxLength Schema Component

Schema Component: [maxLength](https://www.w3.org/TR/xmlschema-2/#dt-maxLength)

{value}

A [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#maxLength-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) other than [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value).

##### 4.3.3.2 XML Representation of maxLength Schema Components

The XML representation for a [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) schema component is a [<maxLength>](https://www.w3.org/TR/xmlschema-2/#element-maxLength) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `maxLength` Element Information Item

<maxLength
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</maxLength>

|[maxLength](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#maxLength-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.3.3 maxLength Validation Rules

**Validation Rule: maxLength Valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength), determined as follows:

1 if the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) then

1.1 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [string](https://www.w3.org/TR/xmlschema-2/#string) or [anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI), then the length of the value, as measured in [character](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-character)s [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be less than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value);

1.2 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary) or [base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary), then the length of the value, as measured in octets of the binary data, [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be less than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value);

1.3 if [{primitive type definition}](https://www.w3.org/TR/xmlschema-2/#defn-primitive) is [QName](https://www.w3.org/TR/xmlschema-2/#QName) or [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION), then any [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value) is facet-valid.

2 if the [{variety}](https://www.w3.org/TR/xmlschema-2/#defn-variety) is [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list), then the length of the value, as measured in list items, [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be less than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value)

The use of [·maxLength·](https://www.w3.org/TR/xmlschema-2/#dt-maxLength) on datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [QName](https://www.w3.org/TR/xmlschema-2/#QName) and [NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION) is deprecated. Future versions of this specification may remove this facet for these datatypes.

##### 4.3.3.4 Constraints on maxLength Schema Components

**Schema Component Constraint: maxLength valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value) is greater than the [{value}](https://www.w3.org/TR/xmlschema-2/#maxLength-value) of the parent [maxLength](https://www.w3.org/TR/xmlschema-2/#dc-maxLength).

#### 4.3.4 pattern

[Definition:]   **pattern** is a constraint on the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of a datatype which is achieved by constraining the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) to literals which match a specific pattern. The value of **pattern**  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

[·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values that are denoted by literals which match a specific [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which is a better representation of postal codes in the United States, by limiting strings to those which are matched by a specific [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

<simpleType name='better-us-zipcode'>
  <restriction base='string'>
    <pattern value='[0-9]{5}(-[0-9]{4})?'/>
  </restriction>
</simpleType>

##### 4.3.4.1 The pattern Schema Component

Schema Component: [pattern](https://www.w3.org/TR/xmlschema-2/#dt-pattern)

{value}

A [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

##### 4.3.4.2 XML Representation of pattern Schema Components

The XML representation for a [pattern](https://www.w3.org/TR/xmlschema-2/#dc-pattern) schema component is a [<pattern>](https://www.w3.org/TR/xmlschema-2/#element-pattern) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `pattern` Element Information Item

<pattern
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [string](https://www.w3.org/TR/xmlschema-2/#string)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</pattern>

[{value}](https://www.w3.org/TR/xmlschema-2/#pattern-value) [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a valid [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

|[pattern](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#pattern-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.4.3 Constraints on XML Representation of pattern

**Schema Representation Constraint: Multiple patterns**

If multiple [<pattern>](https://www.w3.org/TR/xmlschema-2/#element-pattern) element information items appear as [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element) of a [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType), the [[value]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)s should be combined as if they appeared in a single [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex) as separate [·branch·](https://www.w3.org/TR/xmlschema-2/#dt-branch)es.

**Note:**  It is a consequence of the schema representation constraint [Multiple patterns (§4.3.4.3)](https://www.w3.org/TR/xmlschema-2/#src-multiple-patterns) and of the rules for [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction) that [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) facets specified on the _same_ step in a type derivation are **OR**ed together, while [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) facets specified on _different_ steps of a type derivation are **AND**ed together.

Thus, to impose two [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) constraints simultaneously, schema authors may either write a single [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) which expresses the intersection of the two [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern)s they wish to impose, or define each [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) on a separate type derivation step.

##### 4.3.4.4 pattern Validation Rules

**Validation Rule: pattern valid**

A literal in a [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) is facet-valid with respect to [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) if:

1 the literal is among the set of character sequences denoted by the [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex) specified in [{value}](https://www.w3.org/TR/xmlschema-2/#pattern-value).

#### 4.3.5 enumeration

[Definition:]   **enumeration** constrains the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to a specified set of values.

**enumeration** does not impose an order relation on the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) it creates; the value of the [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) property of the [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatype remains that of the datatype from which it is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived).

[·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to a specified set of values.

Example

The following example is a datatype definition for a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which limits the values of dates to the three US holidays enumerated. This datatype definition would appear in a schema authored by an "end-user" and shows how to define a datatype by enumerating the values in its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space). The enumerated values must be type-valid literals for the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype).

<simpleType name='holidays'>
    <annotation>
        <documentation>some US holidays</documentation>
    </annotation>
    <restriction base='gMonthDay'>
      <enumeration value='--01-01'>
        <annotation>
            <documentation>New Year's day</documentation>
        </annotation>
      </enumeration>
      <enumeration value='--07-04'>
        <annotation>
            <documentation>4th of July</documentation>
        </annotation>
      </enumeration>
      <enumeration value='--12-25'>
        <annotation>
            <documentation>Christmas</documentation>
        </annotation>
      </enumeration>
    </restriction>
</simpleType>

##### 4.3.5.1 The enumeration Schema Component

Schema Component: [enumeration](https://www.w3.org/TR/xmlschema-2/#dt-enumeration)

{value}

A set of values from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

##### 4.3.5.2 XML Representation of enumeration Schema Components

The XML representation for an [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration) schema component is an [<enumeration>](https://www.w3.org/TR/xmlschema-2/#element-enumeration) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `enumeration` Element Information Item

<enumeration
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [anySimpleType](https://www.w3.org/TR/xmlschema-2/#dt-anySimpleType)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</enumeration>

[{value}](https://www.w3.org/TR/xmlschema-2/#enumeration-value) [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

|[enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#enumeration-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.5.3 Constraints on XML Representation of enumeration

**Schema Representation Constraint: Multiple enumerations**

If multiple [<enumeration>](https://www.w3.org/TR/xmlschema-2/#element-enumeration) element information items appear as [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element) of a [<simpleType>](https://www.w3.org/TR/xmlschema-2/#element-simpleType) the [{value}](https://www.w3.org/TR/xmlschema-2/#enumeration-value) of the [enumeration](https://www.w3.org/TR/xmlschema-2/#dc-enumeration) component should be the set of all such [[value]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)s.

##### 4.3.5.4 enumeration Validation Rules

**Validation Rule: enumeration valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·enumeration·](https://www.w3.org/TR/xmlschema-2/#dt-enumeration) if the value is one of the values specified in [{value}](https://www.w3.org/TR/xmlschema-2/#enumeration-value)

##### 4.3.5.5 Constraints on enumeration Schema Components

**Schema Component Constraint: enumeration valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if any member of [{value}](https://www.w3.org/TR/xmlschema-2/#enumeration-value) is not in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

#### 4.3.6 whiteSpace

[Definition:]   **whiteSpace** constrains the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of types [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string) such that the various behaviors specified in [Attribute Value Normalization](https://www.w3.org/TR/2000/WD-xml-2e-20000814#AVNormalize) in [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML) are realized. The value of **whiteSpace** must be one of {preserve, replace, collapse}.

preserve

No normalization is done, the value is not changed (this is the behavior required by [[XML 1.0 (Second Edition)]](https://www.w3.org/TR/xmlschema-2/#XML) for element content)

replace

All occurrences of #x9 (tab), #xA (line feed) and #xD (carriage return) are replaced with #x20 (space)

collapse

After the processing implied by **replace**, contiguous sequences of #x20's are collapsed to a single #x20, and leading and trailing #x20's are removed.

**Note:**  The notation #xA used here (and elsewhere in this specification) represents the Universal Character Set (UCS) code point `hexadecimal A` (line feed), which is denoted by U+000A. This notation is to be distinguished from `&#xA;`, which is the XML [character reference](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-CharRef) to that same UCS code point.

**whiteSpace** is applicable to all [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) and [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatypes. For all [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatypes other than [string](https://www.w3.org/TR/xmlschema-2/#string) (and types [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction) from it) the value of **whiteSpace** is `collapse` and cannot be changed by a schema author; for [string](https://www.w3.org/TR/xmlschema-2/#string) the value of **whiteSpace** is `preserve`; for any type [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·restriction·](https://www.w3.org/TR/xmlschema-2/#dt-restriction) from [string](https://www.w3.org/TR/xmlschema-2/#string) the value of **whiteSpace** can be any of the three legal values. For all datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) the value of **whiteSpace** is `collapse` and cannot be changed by a schema author. For all datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union)  **whiteSpace** does not apply directly; however, the normalization behavior of [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) types is controlled by the value of **whiteSpace** on that one of the [·memberTypes·](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes) against which the [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) is successfully validated.

**Note:**  For more information on **whiteSpace**, see the discussion on white space normalization in [Schema Component Details](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#components) in [[XML Schema Part 1: Structures]](https://www.w3.org/TR/xmlschema-2/#structural-schemas).

[·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) according to the white space normalization rules.

Example

The following example is the datatype definition for the [token](https://www.w3.org/TR/xmlschema-2/#token) [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in)  [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatype.

<simpleType name='token'>
    <restriction base='normalizedString'>
      <whiteSpace value='collapse'/>
    </restriction>
</simpleType>

##### 4.3.6.1 The whiteSpace Schema Component

Schema Component: [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace)

{value}

One of `{preserve, replace, collapse}`.

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace) other than [{value}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-value).

##### 4.3.6.2 XML Representation of whiteSpace Schema Components

The XML representation for a [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace) schema component is a [<whiteSpace>](https://www.w3.org/TR/xmlschema-2/#element-whiteSpace) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `whiteSpace` Element Information Item

<whiteSpace
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = (collapse | preserve | replace)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</whiteSpace>

|[whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.6.3 whiteSpace Validation Rules

**Note:**  There are no [·Validation Rule·](https://www.w3.org/TR/xmlschema-2/#dt-cvc)s associated [·whiteSpace·](https://www.w3.org/TR/xmlschema-2/#dt-whiteSpace). For more information, see the discussion on white space normalization in [Schema Component Details](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#components) in [[XML Schema Part 1: Structures]](https://www.w3.org/TR/xmlschema-2/#structural-schemas).

##### 4.3.6.4 Constraints on whiteSpace Schema Components

**Schema Component Constraint: whiteSpace valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and any of the following conditions is true:

1 [{value}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-value) is _replace_ or _preserve_ and the [{value}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-value) of the parent [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace) is _collapse_

2 [{value}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-value) is _preserve_ and the [{value}](https://www.w3.org/TR/xmlschema-2/#whiteSpace-value) of the parent [whiteSpace](https://www.w3.org/TR/xmlschema-2/#dc-whiteSpace) is _replace_

#### 4.3.7 maxInclusive

[Definition:]   **maxInclusive** is the [·inclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-upper-bound) of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) for a datatype with the [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) property. The value of **maxInclusive** [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype).

[·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with a specific [·inclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-upper-bound).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which limits values to integers less than or equal to 100, using [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive).

<simpleType name='one-hundred-or-less'>
  <restriction base='integer'>
    <maxInclusive value='100'/>
  </restriction>
</simpleType>

##### 4.3.7.1 The maxInclusive Schema Component

Schema Component: [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive)

{value}

A value from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) other than [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value).

##### 4.3.7.2 XML Representation of maxInclusive Schema Components

The XML representation for a [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) schema component is a [<maxInclusive>](https://www.w3.org/TR/xmlschema-2/#element-maxInclusive) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `maxInclusive` Element Information Item

<maxInclusive
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [anySimpleType](https://www.w3.org/TR/xmlschema-2/#dt-anySimpleType)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</maxInclusive>

[{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

|[maxInclusive](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false, if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.7.3 maxInclusive Validation Rules

**Validation Rule: maxInclusive Valid**

A value in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive), determined as follows:

1 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _true_, then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be numerically less than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value);

2 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _false_ (i.e., [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is one of the date and time related datatypes), then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be chronologically less than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value);

##### 4.3.7.4 Constraints on maxInclusive Schema Components

**Schema Component Constraint: minInclusive <= maxInclusive**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for the value specified for [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be greater than the value specified for [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) for the same datatype.

**Schema Component Constraint: maxInclusive valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if any of the following conditions is true:

1 [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) is greater than the [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) of the parent [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive)

2 [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) is greater than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) of the parent [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive)

3 [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) is less than the [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) of the parent [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive)

4 [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) is less than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) of the parent [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)

#### 4.3.8 maxExclusive

[Definition:]   **maxExclusive** is the [·exclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-upper-bound) of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) for a datatype with the [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) property. The value of **maxExclusive**  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) or be equal to [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) in [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

[·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with a specific [·exclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-upper-bound).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which limits values to integers less than or equal to 100, using [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive).

<simpleType name='less-than-one-hundred-and-one'>
  <restriction base='integer'>
    <maxExclusive value='101'/>
  </restriction>
</simpleType>

Note that the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of this datatype is identical to the previous one (named 'one-hundred-or-less').

##### 4.3.8.1 The maxExclusive Schema Component

Schema Component: [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive)

{value}

A value from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) other than [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value).

##### 4.3.8.2 XML Representation of maxExclusive Schema Components

The XML representation for a [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) schema component is a [<maxExclusive>](https://www.w3.org/TR/xmlschema-2/#element-maxExclusive) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `maxExclusive` Element Information Item

<maxExclusive
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [anySimpleType](https://www.w3.org/TR/xmlschema-2/#dt-anySimpleType)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</maxExclusive>

[{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

|[maxExclusive](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.8.3 maxExclusive Validation Rules

**Validation Rule: maxExclusive Valid**

A value in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive), determined as follows:

1 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _true_, then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be numerically less than [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value);

2 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _false_ (i.e., [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is one of the date and time related datatypes), then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be chronologically less than [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value);

##### 4.3.8.4 Constraints on maxExclusive Schema Components

**Schema Component Constraint: maxInclusive and maxExclusive**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for both [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) and [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) to be specified in the same derivation step of a datatype definition.

**Schema Component Constraint: minExclusive <= maxExclusive**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for the value specified for [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) to be greater than the value specified for [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) for the same datatype.

**Schema Component Constraint: maxExclusive valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if any of the following conditions is true:

1 [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) is greater than the [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) of the parent [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive)

2 [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) is greater than the [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) of the parent [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive)

3 [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) is less than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) of the parent [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive)

4 [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) is less than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) of the parent [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)

#### 4.3.9 minExclusive

[Definition:]   **minExclusive** is the [·exclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-lower-bound) of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) for a datatype with the [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) property. The value of **minExclusive** [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype) or be equal to [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) in [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

[·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with a specific [·exclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-lower-bound).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which limits values to integers greater than or equal to 100, using [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive).

<simpleType name='more-than-ninety-nine'>
  <restriction base='integer'>
    <minExclusive value='99'/>
  </restriction>
</simpleType>

Note that the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of this datatype is identical to the previous one (named 'one-hundred-or-more').

##### 4.3.9.1 The minExclusive Schema Component

Schema Component: [minExclusive](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive)

{value}

A value from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#minExclusive-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive) other than [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value).

##### 4.3.9.2 XML Representation of minExclusive Schema Components

The XML representation for a [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive) schema component is a [<minExclusive>](https://www.w3.org/TR/xmlschema-2/#element-minExclusive) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `minExclusive` Element Information Item

<minExclusive
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [anySimpleType](https://www.w3.org/TR/xmlschema-2/#dt-anySimpleType)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</minExclusive>

[{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

|[minExclusive](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#minExclusive-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.9.3 minExclusive Validation Rules

**Validation Rule: minExclusive Valid**

A value in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) if:

1 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _true_, then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be numerically greater than [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value);

2 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _false_ (i.e., [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is one of the date and time related datatypes), then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be chronologically greater than [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value);

##### 4.3.9.4 Constraints on minExclusive Schema Components

**Schema Component Constraint: minInclusive and minExclusive**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for both [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) and [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) to be specified for the same datatype.

**Schema Component Constraint: minExclusive < maxInclusive**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for the value specified for [·minExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minExclusive) to be greater than or equal to the value specified for [·maxInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxInclusive) for the same datatype.

**Schema Component Constraint: minExclusive valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if any of the following conditions is true:

1 [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) is less than the [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) of the parent [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)

2 [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) is greater the [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) of the parent [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive)

3 [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) is less than the [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) of the parent [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive)

4 [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) is greater than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) of the parent [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive)

#### 4.3.10 minInclusive

[Definition:]   **minInclusive** is the [·inclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-lower-bound) of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) for a datatype with the [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) property. The value of **minInclusive**  [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype).

[·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) provides for:

- Constraining a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to values with a specific [·inclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-lower-bound).

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which limits values to integers greater than or equal to 100, using [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive).

<simpleType name='one-hundred-or-more'>
  <restriction base='integer'>
    <minInclusive value='100'/>
  </restriction>
</simpleType>

##### 4.3.10.1 The minInclusive Schema Component

Schema Component: [minInclusive](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive)

{value}

A value from the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#minInclusive-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive) other than [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value).

##### 4.3.10.2 XML Representation of minInclusive Schema Components

The XML representation for a [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive) schema component is a [<minInclusive>](https://www.w3.org/TR/xmlschema-2/#element-minInclusive) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `minInclusive` Element Information Item

<minInclusive
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [anySimpleType](https://www.w3.org/TR/xmlschema-2/#dt-anySimpleType)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</minInclusive>

[{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype).

|[minInclusive](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#minInclusive-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.10.3 minInclusive Validation Rules

**Validation Rule: minInclusive Valid**

A value in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered) [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) if:

1 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _true_, then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be numerically greater than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value);

2 if the [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) property in [{fundamental facets}](https://www.w3.org/TR/xmlschema-2/#defn-fund-facets) is _false_ (i.e., [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) is one of the date and time related datatypes), then the value [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be chronologically greater than or equal to [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value);

##### 4.3.10.4 Constraints on minInclusive Schema Components

**Schema Component Constraint: minInclusive < maxExclusive**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for the value specified for [·minInclusive·](https://www.w3.org/TR/xmlschema-2/#dt-minInclusive) to be greater than or equal to the value specified for [·maxExclusive·](https://www.w3.org/TR/xmlschema-2/#dt-maxExclusive) for the same datatype.

**Schema Component Constraint: minInclusive valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if any of the following conditions is true:

1 [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) is less than the [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) of the parent [minInclusive](https://www.w3.org/TR/xmlschema-2/#dc-minInclusive)

2 [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) is greater the [{value}](https://www.w3.org/TR/xmlschema-2/#maxInclusive-value) of the parent [maxInclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxInclusive)

3 [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) is less than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#minExclusive-value) of the parent [minExclusive](https://www.w3.org/TR/xmlschema-2/#dc-minExclusive)

4 [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#minInclusive-value) is greater than or equal to the [{value}](https://www.w3.org/TR/xmlschema-2/#maxExclusive-value) of the parent [maxExclusive](https://www.w3.org/TR/xmlschema-2/#dc-maxExclusive)

#### 4.3.11 totalDigits

[Definition:]   **totalDigits** controls the maximum number of values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [decimal](https://www.w3.org/TR/xmlschema-2/#decimal), by restricting it to numbers that are expressible as _i × 10^-n_ where _i_ and _n_ are integers such that _|i| < 10^totalDigits_ and _0 <= n <= totalDigits_. The value of **totalDigits** [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a [positiveInteger](https://www.w3.org/TR/xmlschema-2/#positiveInteger).

The term **totalDigits** is chosen to reflect the fact that it restricts the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to those values that can be represented lexically using at most _totalDigits_ digits. Note that it does not restrict the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) directly; a lexical representation that adds additional leading zero digits or trailing fractional zero digits is still permitted.

##### 4.3.11.1 The totalDigits Schema Component

Schema Component: [totalDigits](https://www.w3.org/TR/xmlschema-2/#dt-totalDigits)

{value}

A [positiveInteger](https://www.w3.org/TR/xmlschema-2/#positiveInteger).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#totalDigits-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [totalDigits](https://www.w3.org/TR/xmlschema-2/#dc-totalDigits) other than [{value}](https://www.w3.org/TR/xmlschema-2/#totalDigits-value).

##### 4.3.11.2 XML Representation of totalDigits Schema Components

The XML representation for a [totalDigits](https://www.w3.org/TR/xmlschema-2/#dc-totalDigits) schema component is a [<totalDigits>](https://www.w3.org/TR/xmlschema-2/#element-totalDigits) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `totalDigits` Element Information Item

<totalDigits
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [positiveInteger](https://www.w3.org/TR/xmlschema-2/#positiveInteger)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</totalDigits>

|[totalDigits](https://www.w3.org/TR/xmlschema-2/#dc-totalDigits) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#totalDigits-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#totalDigits-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.11.3 totalDigits Validation Rules

**Validation Rule: totalDigits Valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·totalDigits·](https://www.w3.org/TR/xmlschema-2/#dt-totalDigits) if:

1 that value is expressible as _i × 10^-n_ where _i_ and _n_ are integers such that _|i| < 10^[{value}](https://www.w3.org/TR/xmlschema-2/#totalDigits-value)_ and _0 <= n <= [{value}](https://www.w3.org/TR/xmlschema-2/#totalDigits-value)_.

##### 4.3.11.4 Constraints on totalDigits Schema Components

**Schema Component Constraint: totalDigits valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [totalDigits](https://www.w3.org/TR/xmlschema-2/#dc-totalDigits) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#totalDigits-value) is greater than the [{value}](https://www.w3.org/TR/xmlschema-2/#totalDigits-value) of the parent [totalDigits](https://www.w3.org/TR/xmlschema-2/#dc-totalDigits)

#### 4.3.12 fractionDigits

[Definition:]   **fractionDigits** controls the size of the minimum difference between values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) of datatypes [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from **decimal**, by restricting the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to numbers that are expressible as _i × 10^-n_ where _i_ and _n_ are integers and _0 <= n <= fractionDigits_. The value of **fractionDigits** [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) be a [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

The term **fractionDigits** is chosen to reflect the fact that it restricts the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) to those values that can be represented lexically using at most _fractionDigits_ to the right of the decimal point. Note that it does not restrict the [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) directly; a non-[·canonical lexical representation·](https://www.w3.org/TR/xmlschema-2/#dt-canonical-representation) that adds additional leading zero digits or trailing fractional zero digits is still permitted.

Example

The following is the definition of a [·user-derived·](https://www.w3.org/TR/xmlschema-2/#dt-user-derived) datatype which could be used to represent the magnitude of a person's body temperature on the Celsius scale. This definition would appear in a schema authored by an "end-user" and shows how to define a datatype by specifying facet values which constrain the range of the [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype).

<simpleType name='celsiusBodyTemp'>
  <restriction base='decimal'>
    <totalDigits value='4'/>
    <fractionDigits value='1'/>
    <minInclusive value='36.4'/>
    <maxInclusive value='40.5'/>
  </restriction>
</simpleType>

##### 4.3.12.1 The fractionDigits Schema Component

Schema Component: [fractionDigits](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits)

{value}

A [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger).

{fixed}

A [boolean](https://www.w3.org/TR/xmlschema-2/#boolean).

{annotation}

Optional. An [annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#Annotation).

If [{fixed}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-fixed) is _true_, then types for which the current type is the [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) cannot specify a value for [fractionDigits](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) other than [{value}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-value).

##### 4.3.12.2 XML Representation of fractionDigits Schema Components

The XML representation for a [fractionDigits](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) schema component is a [<fractionDigits>](https://www.w3.org/TR/xmlschema-2/#element-fractionDigits) element information item. The correspondences between the properties of the information item and properties of the component are as follows:

XML Representation Summary: `fractionDigits` Element Information Item

<fractionDigits
  fixed = [boolean](https://www.w3.org/TR/xmlschema-2/#boolean) : false
  id = [ID](https://www.w3.org/TR/xmlschema-2/#ID)
  **value** = [nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)
  _{any attributes with non-schema namespace . . .}_>
  _Content:_ ([annotation](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation)?)
</fractionDigits>

|[fractionDigits](https://www.w3.org/TR/xmlschema-2/#dc-fractionDigits) **Schema Component**|
|---|
|\|Property\|Representation\|<br>\|:--\|:--\|<br>\|[{value}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-value)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `value` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute)\|<br>\|[{fixed}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-fixed)\|The [actual value](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-vv) of the `fixed` [[attribute]](https://www.w3.org/TR/xml-infoset/#infoitem.attribute), if present, otherwise false\|<br>\|[{annotation}](https://www.w3.org/TR/xmlschema-2/#defn-annotation)\|The annotations corresponding to all the [<annotation>](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#element-annotation) element information items in the [[children]](https://www.w3.org/TR/xml-infoset/#infoitem.element), if any.\||

##### 4.3.12.3 fractionDigits Validation Rules

**Validation Rule: fractionDigits Valid**

A value in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is facet-valid with respect to [·fractionDigits·](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits) if:

1 that value is expressible as _i × 10^-n_ where _i_ and _n_ are integers and _0 <= n <= [{value}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-value)_.

##### 4.3.12.4 Constraints on fractionDigits Schema Components

**Schema Component Constraint: fractionDigits less than or equal to totalDigits**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) for [·fractionDigits·](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits) to be greater than [·totalDigits·](https://www.w3.org/TR/xmlschema-2/#dt-totalDigits).

**Schema Component Constraint: fractionDigits valid restriction**

It is an [·error·](https://www.w3.org/TR/xmlschema-2/#dt-error) if [·fractionDigits·](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits) is among the members of [{facets}](https://www.w3.org/TR/xmlschema-2/#defn-facets) of [{base type definition}](https://www.w3.org/TR/xmlschema-2/#defn-basetype) and [{value}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-value) is greater than the [{value}](https://www.w3.org/TR/xmlschema-2/#fractionDigits-value) of the parent [·fractionDigits·](https://www.w3.org/TR/xmlschema-2/#dt-fractionDigits).

## 5 Conformance

This specification describes two levels of conformance for datatype processors. The first is required of all processors. Support for the other will depend on the application environments for which the processor is intended.

[Definition:]   **Minimally conforming** processors [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) completely and correctly implement the [·Constraint on Schemas·](https://www.w3.org/TR/xmlschema-2/#dt-cos) and [·Validation Rule·](https://www.w3.org/TR/xmlschema-2/#dt-cvc) .

[Definition:]   Processors which accept schemas in the form of XML documents as described in [XML Representation of Simple Type Definition Schema Components (§4.1.2)](https://www.w3.org/TR/xmlschema-2/#xr-defn) (and other relevant portions of [Datatype components (§4)](https://www.w3.org/TR/xmlschema-2/#datatype-components)) are additionally said to provide **conformance to the XML Representation of Schemas**, and [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must), when processing schema documents, completely and correctly implement all [·Schema Representation Constraint·](https://www.w3.org/TR/xmlschema-2/#dt-src)s in this specification, and [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) adhere exactly to the specifications in [XML Representation of Simple Type Definition Schema Components (§4.1.2)](https://www.w3.org/TR/xmlschema-2/#xr-defn) (and other relevant portions of [Datatype components (§4)](https://www.w3.org/TR/xmlschema-2/#datatype-components)) for mapping the contents of such documents to [schema components](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-component) for use in validation.

**Note:**  By separating the conformance requirements relating to the concrete syntax of XML schema documents, this specification admits processors which validate using schemas stored in optimized binary representations, dynamically created schemas represented as programming language data structures, or implementations in which particular schemas are compiled into executable code such as C or Java. Such processors can be said to be [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) but not necessarily in [·conformance to the XML Representation of Schemas·](https://www.w3.org/TR/xmlschema-2/#dt-interchange).

## A Schema for Datatype Definitions (normative)

<!DOCTYPE xs:schema PUBLIC "-//W3C//DTD XMLSCHEMA 200102//EN" "XMLSchema.dtd" [

<!--
     keep this schema XML1.0 DTD valid
  -->
        <!ENTITY % schemaAttrs 'xmlns:hfp CDATA #IMPLIED'>

        <!ELEMENT hfp:hasFacet EMPTY>
        <!ATTLIST hfp:hasFacet
                name NMTOKEN #REQUIRED>

        <!ELEMENT hfp:hasProperty EMPTY>
        <!ATTLIST hfp:hasProperty
                name NMTOKEN #REQUIRED
                value CDATA #REQUIRED>
<!--
        Make sure that processors that do not read the external
        subset will know about the various IDs we declare
  -->
        <!ATTLIST xs:simpleType id ID #IMPLIED>
        <!ATTLIST xs:maxExclusive id ID #IMPLIED>
        <!ATTLIST xs:minExclusive id ID #IMPLIED>
        <!ATTLIST xs:maxInclusive id ID #IMPLIED>
        <!ATTLIST xs:minInclusive id ID #IMPLIED>
        <!ATTLIST xs:totalDigits id ID #IMPLIED>
        <!ATTLIST xs:fractionDigits id ID #IMPLIED>
        <!ATTLIST xs:length id ID #IMPLIED>
        <!ATTLIST xs:minLength id ID #IMPLIED>
        <!ATTLIST xs:maxLength id ID #IMPLIED>
        <!ATTLIST xs:enumeration id ID #IMPLIED>
        <!ATTLIST xs:pattern id ID #IMPLIED>
        <!ATTLIST xs:appinfo id ID #IMPLIED>
        <!ATTLIST xs:documentation id ID #IMPLIED>
        <!ATTLIST xs:list id ID #IMPLIED>
        <!ATTLIST xs:union id ID #IMPLIED>
        ]>

<?xml version='1.0'?>
<xs:schema xmlns:hfp="http://www.w3.org/2001/XMLSchema-hasFacetAndProperty"
           xmlns:xs="http://www.w3.org/2001/XMLSchema" blockDefault="#all"
           elementFormDefault="qualified" xml:lang="en"
           targetNamespace="http://www.w3.org/2001/XMLSchema"
           version="Id: datatypes.xsd,v 1.4 2004/05/29 10:26:33 ht Exp ">
  <xs:annotation>
    <xs:documentation source="../datatypes/datatypes-with-errata.html">
      The schema corresponding to this document is normative,
      with respect to the syntactic constraints it expresses in the
      XML Schema language.  The documentation (within &lt;documentation>
      elements) below, is not normative, but rather highlights important
      aspects of the W3C Recommendation of which this is a part
    </xs:documentation>
  </xs:annotation>
  <xs:annotation>
    <xs:documentation>
      First the built-in primitive datatypes.  These definitions are for
      information only, the real built-in definitions are magic.
    </xs:documentation>
    <xs:documentation>
      For each built-in datatype in this schema (both primitive and
      derived) can be uniquely addressed via a URI constructed
      as follows:
        1) the base URI is the URI of the XML Schema namespace
        2) the fragment identifier is the name of the datatype

      For example, to address the int datatype, the URI is:

        http://www.w3.org/2001/XMLSchema#int

      Additionally, each facet definition element can be uniquely
      addressed via a URI constructed as follows:
        1) the base URI is the URI of the XML Schema namespace
        2) the fragment identifier is the name of the facet

      For example, to address the maxInclusive facet, the URI is:

        http://www.w3.org/2001/XMLSchema#maxInclusive

      Additionally, each facet usage in a built-in datatype definition
      can be uniquely addressed via a URI constructed as follows:
        1) the base URI is the URI of the XML Schema namespace
        2) the fragment identifier is the name of the datatype, followed
           by a period (".") followed by the name of the facet

      For example, to address the usage of the maxInclusive facet in
      the definition of int, the URI is:

        http://www.w3.org/2001/XMLSchema#int.maxInclusive

    </xs:documentation>
  </xs:annotation>
  <xs:simpleType name="string" id="string">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#string"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace value="preserve" id="string.preserve"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="boolean" id="boolean">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="finite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#boolean"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="boolean.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="float" id="float">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="true"/>
        <hfp:hasProperty name="cardinality" value="finite"/>
        <hfp:hasProperty name="numeric" value="true"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#float"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="float.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="double" id="double">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="true"/>
        <hfp:hasProperty name="cardinality" value="finite"/>
        <hfp:hasProperty name="numeric" value="true"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#double"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="double.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="decimal" id="decimal">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="totalDigits"/>
        <hfp:hasFacet name="fractionDigits"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="total"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="true"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#decimal"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="decimal.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="duration" id="duration">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#duration"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="duration.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="dateTime" id="dateTime">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#dateTime"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="dateTime.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="time" id="time">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#time"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="time.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="date" id="date">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#date"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="date.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="gYearMonth" id="gYearMonth">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#gYearMonth"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="gYearMonth.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="gYear" id="gYear">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#gYear"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="gYear.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="gMonthDay" id="gMonthDay">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#gMonthDay"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="gMonthDay.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="gDay" id="gDay">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#gDay"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="gDay.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="gMonth" id="gMonth">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="maxInclusive"/>
        <hfp:hasFacet name="maxExclusive"/>
        <hfp:hasFacet name="minInclusive"/>
        <hfp:hasFacet name="minExclusive"/>
        <hfp:hasProperty name="ordered" value="partial"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#gMonth"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="gMonth.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="hexBinary" id="hexBinary">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#binary"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="hexBinary.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="base64Binary" id="base64Binary">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#base64Binary"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="base64Binary.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="anyURI" id="anyURI">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#anyURI"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="anyURI.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="QName" id="QName">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#QName"/>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="QName.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="NOTATION" id="NOTATION">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#NOTATION"/>
      <xs:documentation>
        NOTATION cannot be used directly in a schema; rather a type
        must be derived from it by specifying at least one enumeration
        facet whose value is the name of a NOTATION declared in the
        schema.
      </xs:documentation>
    </xs:annotation>
    <xs:restriction base="xs:anySimpleType">
      <xs:whiteSpace fixed="true" value="collapse" id="NOTATION.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:annotation>
    <xs:documentation>
      Now the derived primitive types
    </xs:documentation>
  </xs:annotation>
  <xs:simpleType name="normalizedString" id="normalizedString">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#normalizedString"/>
    </xs:annotation>
    <xs:restriction base="xs:string">
      <xs:whiteSpace value="replace" id="normalizedString.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="token" id="token">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#token"/>
    </xs:annotation>
    <xs:restriction base="xs:normalizedString">
      <xs:whiteSpace value="collapse" id="token.whiteSpace"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="language" id="language">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#language"/>
    </xs:annotation>
    <xs:restriction base="xs:token">
      <xs:pattern value="[a-zA-Z]{1,8}(-[a-zA-Z0-9]{1,8})*"
                  id="language.pattern">
        <xs:annotation>
          <xs:documentation source="http://www.ietf.org/rfc/rfc3066.txt">
            pattern specifies the content of section 2.12 of XML 1.0e2
            and RFC 3066 (Revised version of RFC 1766).
          </xs:documentation>
        </xs:annotation>
      </xs:pattern>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="IDREFS" id="IDREFS">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#IDREFS"/>
    </xs:annotation>
    <xs:restriction>
      <xs:simpleType>
        <xs:list itemType="xs:IDREF"/>
      </xs:simpleType>
      <xs:minLength value="1" id="IDREFS.minLength"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="ENTITIES" id="ENTITIES">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#ENTITIES"/>
    </xs:annotation>
    <xs:restriction>
      <xs:simpleType>
        <xs:list itemType="xs:ENTITY"/>
      </xs:simpleType>
      <xs:minLength value="1" id="ENTITIES.minLength"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="NMTOKEN" id="NMTOKEN">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#NMTOKEN"/>
    </xs:annotation>
    <xs:restriction base="xs:token">
      <xs:pattern value="\c+" id="NMTOKEN.pattern">
        <xs:annotation>
          <xs:documentation source="http://www.w3.org/TR/REC-xml#NT-Nmtoken">
            pattern matches production 7 from the XML spec
          </xs:documentation>
        </xs:annotation>
      </xs:pattern>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="NMTOKENS" id="NMTOKENS">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasFacet name="length"/>
        <hfp:hasFacet name="minLength"/>
        <hfp:hasFacet name="maxLength"/>
        <hfp:hasFacet name="enumeration"/>
        <hfp:hasFacet name="whiteSpace"/>
        <hfp:hasFacet name="pattern"/>
        <hfp:hasProperty name="ordered" value="false"/>
        <hfp:hasProperty name="bounded" value="false"/>
        <hfp:hasProperty name="cardinality" value="countably infinite"/>
        <hfp:hasProperty name="numeric" value="false"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#NMTOKENS"/>
    </xs:annotation>
    <xs:restriction>
      <xs:simpleType>
        <xs:list itemType="xs:NMTOKEN"/>
      </xs:simpleType>
      <xs:minLength value="1" id="NMTOKENS.minLength"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="Name" id="Name">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#Name"/>
    </xs:annotation>
    <xs:restriction base="xs:token">
      <xs:pattern value="\i\c*" id="Name.pattern">
        <xs:annotation>
          <xs:documentation source="http://www.w3.org/TR/REC-xml#NT-Name">
            pattern matches production 5 from the XML spec
          </xs:documentation>
        </xs:annotation>
      </xs:pattern>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="NCName" id="NCName">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#NCName"/>
    </xs:annotation>
    <xs:restriction base="xs:Name">
      <xs:pattern value="[\i-[:]][\c-[:]]*" id="NCName.pattern">
        <xs:annotation>
          <xs:documentation
               source="http://www.w3.org/TR/REC-xml-names/#NT-NCName">
            pattern matches production 4 from the Namespaces in XML spec
          </xs:documentation>
        </xs:annotation>
      </xs:pattern>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="ID" id="ID">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#ID"/>
    </xs:annotation>
    <xs:restriction base="xs:NCName"/>
  </xs:simpleType>
  <xs:simpleType name="IDREF" id="IDREF">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#IDREF"/>
    </xs:annotation>
    <xs:restriction base="xs:NCName"/>
  </xs:simpleType>
  <xs:simpleType name="ENTITY" id="ENTITY">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#ENTITY"/>
    </xs:annotation>
    <xs:restriction base="xs:NCName"/>
  </xs:simpleType>
  <xs:simpleType name="integer" id="integer">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#integer"/>
    </xs:annotation>
    <xs:restriction base="xs:decimal">
      <xs:fractionDigits fixed="true" value="0" id="integer.fractionDigits"/>
      <xs:pattern value="[\-+]?[0-9]+"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="nonPositiveInteger" id="nonPositiveInteger">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#nonPositiveInteger"/>
    </xs:annotation>
    <xs:restriction base="xs:integer">
      <xs:maxInclusive value="0" id="nonPositiveInteger.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="negativeInteger" id="negativeInteger">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#negativeInteger"/>
    </xs:annotation>
    <xs:restriction base="xs:nonPositiveInteger">
      <xs:maxInclusive value="-1" id="negativeInteger.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="long" id="long">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasProperty name="bounded" value="true"/>
        <hfp:hasProperty name="cardinality" value="finite"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#long"/>
    </xs:annotation>
    <xs:restriction base="xs:integer">
      <xs:minInclusive value="-9223372036854775808" id="long.minInclusive"/>
      <xs:maxInclusive value="9223372036854775807" id="long.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="int" id="int">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#int"/>
    </xs:annotation>
    <xs:restriction base="xs:long">
      <xs:minInclusive value="-2147483648" id="int.minInclusive"/>
      <xs:maxInclusive value="2147483647" id="int.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="short" id="short">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#short"/>
    </xs:annotation>
    <xs:restriction base="xs:int">
      <xs:minInclusive value="-32768" id="short.minInclusive"/>
      <xs:maxInclusive value="32767" id="short.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="byte" id="byte">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#byte"/>
    </xs:annotation>
    <xs:restriction base="xs:short">
      <xs:minInclusive value="-128" id="byte.minInclusive"/>
      <xs:maxInclusive value="127" id="byte.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="nonNegativeInteger" id="nonNegativeInteger">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#nonNegativeInteger"/>
    </xs:annotation>
    <xs:restriction base="xs:integer">
      <xs:minInclusive value="0" id="nonNegativeInteger.minInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="unsignedLong" id="unsignedLong">
    <xs:annotation>
      <xs:appinfo>
        <hfp:hasProperty name="bounded" value="true"/>
        <hfp:hasProperty name="cardinality" value="finite"/>
      </xs:appinfo>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#unsignedLong"/>
    </xs:annotation>
    <xs:restriction base="xs:nonNegativeInteger">
      <xs:maxInclusive value="18446744073709551615"
                       id="unsignedLong.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="unsignedInt" id="unsignedInt">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#unsignedInt"/>
    </xs:annotation>
    <xs:restriction base="xs:unsignedLong">
      <xs:maxInclusive value="4294967295" id="unsignedInt.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="unsignedShort" id="unsignedShort">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#unsignedShort"/>
    </xs:annotation>
    <xs:restriction base="xs:unsignedInt">
      <xs:maxInclusive value="65535" id="unsignedShort.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="unsignedByte" id="unsignedByte">
    <xs:annotation>
      <xs:documentation source="http://www.w3.org/TR/xmlschema-2/#unsignedByte"/>
    </xs:annotation>
    <xs:restriction base="xs:unsignedShort">
      <xs:maxInclusive value="255" id="unsignedByte.maxInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="positiveInteger" id="positiveInteger">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#positiveInteger"/>
    </xs:annotation>
    <xs:restriction base="xs:nonNegativeInteger">
      <xs:minInclusive value="1" id="positiveInteger.minInclusive"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:simpleType name="derivationControl">
    <xs:annotation>
      <xs:documentation>
   A utility type, not for public use</xs:documentation>
    </xs:annotation>
    <xs:restriction base="xs:NMTOKEN">
      <xs:enumeration value="substitution"/>
      <xs:enumeration value="extension"/>
      <xs:enumeration value="restriction"/>
      <xs:enumeration value="list"/>
      <xs:enumeration value="union"/>
    </xs:restriction>
  </xs:simpleType>
  <xs:group name="simpleDerivation">
    <xs:choice>
      <xs:element ref="xs:restriction"/>
      <xs:element ref="xs:list"/>
      <xs:element ref="xs:union"/>
    </xs:choice>
  </xs:group>
  <xs:simpleType name="simpleDerivationSet">
    <xs:annotation>
      <xs:documentation>
   #all or (possibly empty) subset of {restriction, union, list}
   </xs:documentation>
      <xs:documentation>
   A utility type, not for public use</xs:documentation>
    </xs:annotation>
    <xs:union>
      <xs:simpleType>
        <xs:restriction base="xs:token">
          <xs:enumeration value="#all"/>
        </xs:restriction>
      </xs:simpleType>
      <xs:simpleType>
        <xs:list>
          <xs:simpleType>
            <xs:restriction base="xs:derivationControl">
              <xs:enumeration value="list"/>
              <xs:enumeration value="union"/>
              <xs:enumeration value="restriction"/>
            </xs:restriction>
          </xs:simpleType>
        </xs:list>
      </xs:simpleType>
    </xs:union>
  </xs:simpleType>
  <xs:complexType name="simpleType" abstract="true">
    <xs:complexContent>
      <xs:extension base="xs:annotated">
        <xs:group ref="xs:simpleDerivation"/>
        <xs:attribute name="final" type="xs:simpleDerivationSet"/>
        <xs:attribute name="name" type="xs:NCName">
          <xs:annotation>
            <xs:documentation>
              Can be restricted to required or forbidden
            </xs:documentation>
          </xs:annotation>
        </xs:attribute>
      </xs:extension>
    </xs:complexContent>
  </xs:complexType>
  <xs:complexType name="topLevelSimpleType">
    <xs:complexContent>
      <xs:restriction base="xs:simpleType">
        <xs:sequence>
          <xs:element ref="xs:annotation" minOccurs="0"/>
          <xs:group ref="xs:simpleDerivation"/>
        </xs:sequence>
        <xs:attribute name="name" type="xs:NCName" use="required">
          <xs:annotation>
            <xs:documentation>
              Required at the top level
            </xs:documentation>
          </xs:annotation>
        </xs:attribute>
        <xs:anyAttribute namespace="##other" processContents="lax"/>
      </xs:restriction>
    </xs:complexContent>
  </xs:complexType>
  <xs:complexType name="localSimpleType">
    <xs:complexContent>
      <xs:restriction base="xs:simpleType">
        <xs:sequence>
          <xs:element ref="xs:annotation" minOccurs="0"/>
          <xs:group ref="xs:simpleDerivation"/>
        </xs:sequence>
        <xs:attribute name="name" use="prohibited">
          <xs:annotation>
            <xs:documentation>
              Forbidden when nested
            </xs:documentation>
          </xs:annotation>
        </xs:attribute>
        <xs:attribute name="final" use="prohibited"/>
        <xs:anyAttribute namespace="##other" processContents="lax"/>
      </xs:restriction>
    </xs:complexContent>
  </xs:complexType>
  <xs:element name="simpleType" type="xs:topLevelSimpleType" id="simpleType">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-simpleType"/>
    </xs:annotation>
  </xs:element>
  <xs:group name="facets">
    <xs:annotation>
      <xs:documentation>
       We should use a substitution group for facets, but
       that's ruled out because it would allow users to
       add their own, which we're not ready for yet.
    </xs:documentation>
    </xs:annotation>
    <xs:choice>
      <xs:element ref="xs:minExclusive"/>
      <xs:element ref="xs:minInclusive"/>
      <xs:element ref="xs:maxExclusive"/>
      <xs:element ref="xs:maxInclusive"/>
      <xs:element ref="xs:totalDigits"/>
      <xs:element ref="xs:fractionDigits"/>
      <xs:element ref="xs:length"/>
      <xs:element ref="xs:minLength"/>
      <xs:element ref="xs:maxLength"/>
      <xs:element ref="xs:enumeration"/>
      <xs:element ref="xs:whiteSpace"/>
      <xs:element ref="xs:pattern"/>
    </xs:choice>
  </xs:group>
  <xs:group name="simpleRestrictionModel">
    <xs:sequence>
      <xs:element name="simpleType" type="xs:localSimpleType" minOccurs="0"/>
      <xs:group ref="xs:facets" minOccurs="0" maxOccurs="unbounded"/>
    </xs:sequence>
  </xs:group>
  <xs:element name="restriction" id="restriction">
    <xs:complexType>
      <xs:annotation>
        <xs:documentation
             source="http://www.w3.org/TR/xmlschema-2/#element-restriction">
          base attribute and simpleType child are mutually
          exclusive, but one or other is required
        </xs:documentation>
      </xs:annotation>
      <xs:complexContent>
        <xs:extension base="xs:annotated">
          <xs:group ref="xs:simpleRestrictionModel"/>
          <xs:attribute name="base" type="xs:QName" use="optional"/>
        </xs:extension>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
  <xs:element name="list" id="list">
    <xs:complexType>
      <xs:annotation>
        <xs:documentation
             source="http://www.w3.org/TR/xmlschema-2/#element-list">
          itemType attribute and simpleType child are mutually
          exclusive, but one or other is required
        </xs:documentation>
      </xs:annotation>
      <xs:complexContent>
        <xs:extension base="xs:annotated">
          <xs:sequence>
            <xs:element name="simpleType" type="xs:localSimpleType"
                        minOccurs="0"/>
          </xs:sequence>
          <xs:attribute name="itemType" type="xs:QName" use="optional"/>
        </xs:extension>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
  <xs:element name="union" id="union">
    <xs:complexType>
      <xs:annotation>
        <xs:documentation
             source="http://www.w3.org/TR/xmlschema-2/#element-union">
          memberTypes attribute must be non-empty or there must be
          at least one simpleType child
        </xs:documentation>
      </xs:annotation>
      <xs:complexContent>
        <xs:extension base="xs:annotated">
          <xs:sequence>
            <xs:element name="simpleType" type="xs:localSimpleType"
                        minOccurs="0" maxOccurs="unbounded"/>
          </xs:sequence>
          <xs:attribute name="memberTypes" use="optional">
            <xs:simpleType>
              <xs:list itemType="xs:QName"/>
            </xs:simpleType>
          </xs:attribute>
        </xs:extension>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
  <xs:complexType name="facet">
    <xs:complexContent>
      <xs:extension base="xs:annotated">
        <xs:attribute name="value" use="required"/>
        <xs:attribute name="fixed" type="xs:boolean" default="false"
                      use="optional"/>
      </xs:extension>
    </xs:complexContent>
  </xs:complexType>
  <xs:complexType name="noFixedFacet">
    <xs:complexContent>
      <xs:restriction base="xs:facet">
        <xs:sequence>
          <xs:element ref="xs:annotation" minOccurs="0"/>
        </xs:sequence>
        <xs:attribute name="fixed" use="prohibited"/>
        <xs:anyAttribute namespace="##other" processContents="lax"/>
      </xs:restriction>
    </xs:complexContent>
  </xs:complexType>
  <xs:element name="minExclusive" type="xs:facet" id="minExclusive">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-minExclusive"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="minInclusive" type="xs:facet" id="minInclusive">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-minInclusive"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="maxExclusive" type="xs:facet" id="maxExclusive">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-maxExclusive"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="maxInclusive" type="xs:facet" id="maxInclusive">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-maxInclusive"/>
    </xs:annotation>
  </xs:element>
  <xs:complexType name="numFacet">
    <xs:complexContent>
      <xs:restriction base="xs:facet">
        <xs:sequence>
          <xs:element ref="xs:annotation" minOccurs="0"/>
        </xs:sequence>
        <xs:attribute name="value" type="xs:nonNegativeInteger" use="required"/>
        <xs:anyAttribute namespace="##other" processContents="lax"/>
      </xs:restriction>
    </xs:complexContent>
  </xs:complexType>
  <xs:element name="totalDigits" id="totalDigits">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-totalDigits"/>
    </xs:annotation>
    <xs:complexType>
      <xs:complexContent>
        <xs:restriction base="xs:numFacet">
          <xs:sequence>
            <xs:element ref="xs:annotation" minOccurs="0"/>
          </xs:sequence>
          <xs:attribute name="value" type="xs:positiveInteger" use="required"/>
          <xs:anyAttribute namespace="##other" processContents="lax"/>
        </xs:restriction>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
  <xs:element name="fractionDigits" type="xs:numFacet" id="fractionDigits">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-fractionDigits"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="length" type="xs:numFacet" id="length">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-length"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="minLength" type="xs:numFacet" id="minLength">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-minLength"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="maxLength" type="xs:numFacet" id="maxLength">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-maxLength"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="enumeration" type="xs:noFixedFacet" id="enumeration">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-enumeration"/>
    </xs:annotation>
  </xs:element>
  <xs:element name="whiteSpace" id="whiteSpace">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-whiteSpace"/>
    </xs:annotation>
    <xs:complexType>
      <xs:complexContent>
        <xs:restriction base="xs:facet">
          <xs:sequence>
            <xs:element ref="xs:annotation" minOccurs="0"/>
          </xs:sequence>
          <xs:attribute name="value" use="required">
            <xs:simpleType>
              <xs:restriction base="xs:NMTOKEN">
                <xs:enumeration value="preserve"/>
                <xs:enumeration value="replace"/>
                <xs:enumeration value="collapse"/>
              </xs:restriction>
            </xs:simpleType>
          </xs:attribute>
          <xs:anyAttribute namespace="##other" processContents="lax"/>
        </xs:restriction>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
  <xs:element name="pattern" id="pattern">
    <xs:annotation>
      <xs:documentation
           source="http://www.w3.org/TR/xmlschema-2/#element-pattern"/>
    </xs:annotation>
    <xs:complexType>
      <xs:complexContent>
        <xs:restriction base="xs:noFixedFacet">
          <xs:sequence>
            <xs:element ref="xs:annotation" minOccurs="0"/>
          </xs:sequence>
          <xs:attribute name="value" type="xs:string" use="required"/>
          <xs:anyAttribute namespace="##other" processContents="lax"/>
        </xs:restriction>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
</xs:schema>

## B DTD for Datatype Definitions (non-normative)

<!--
        DTD for XML Schemas: Part 2: Datatypes
        Id: datatypes.dtd,v  1.1 2003/08/28 13:30:52 ht Exp
        Note this DTD is NOT normative, or even definitive.
  -->

<!--
        This DTD cannot be used on its own, it is intended
        only for incorporation in XMLSchema.dtd, q.v.
  -->

<!-- Define all the element names, with optional prefix -->
<!ENTITY % simpleType "%p;simpleType">
<!ENTITY % restriction "%p;restriction">
<!ENTITY % list "%p;list">
<!ENTITY % union "%p;union">
<!ENTITY % maxExclusive "%p;maxExclusive">
<!ENTITY % minExclusive "%p;minExclusive">
<!ENTITY % maxInclusive "%p;maxInclusive">
<!ENTITY % minInclusive "%p;minInclusive">
<!ENTITY % totalDigits "%p;totalDigits">
<!ENTITY % fractionDigits "%p;fractionDigits">
<!ENTITY % length "%p;length">
<!ENTITY % minLength "%p;minLength">
<!ENTITY % maxLength "%p;maxLength">
<!ENTITY % enumeration "%p;enumeration">
<!ENTITY % whiteSpace "%p;whiteSpace">
<!ENTITY % pattern "%p;pattern">

<!--
        Customisation entities for the ATTLIST of each element
        type. Define one of these if your schema takes advantage
        of the anyAttribute='##other' in the schema for schemas
  -->

<!ENTITY % simpleTypeAttrs "">
<!ENTITY % restrictionAttrs "">
<!ENTITY % listAttrs "">
<!ENTITY % unionAttrs "">
<!ENTITY % maxExclusiveAttrs "">
<!ENTITY % minExclusiveAttrs "">
<!ENTITY % maxInclusiveAttrs "">
<!ENTITY % minInclusiveAttrs "">
<!ENTITY % totalDigitsAttrs "">
<!ENTITY % fractionDigitsAttrs "">
<!ENTITY % lengthAttrs "">
<!ENTITY % minLengthAttrs "">
<!ENTITY % maxLengthAttrs "">
<!ENTITY % enumerationAttrs "">
<!ENTITY % whiteSpaceAttrs "">
<!ENTITY % patternAttrs "">

<!-- Define some entities for informative use as attribute
        types -->
<!ENTITY % URIref "CDATA">
<!ENTITY % XPathExpr "CDATA">
<!ENTITY % QName "NMTOKEN">
<!ENTITY % QNames "NMTOKENS">
<!ENTITY % NCName "NMTOKEN">
<!ENTITY % nonNegativeInteger "NMTOKEN">
<!ENTITY % boolean "(true|false)">
<!ENTITY % simpleDerivationSet "CDATA">
<!--
        #all or space-separated list drawn from derivationChoice
  -->

<!--
        Note that the use of 'facet' below is less restrictive
        than is really intended:  There should in fact be no
        more than one of each of minInclusive, minExclusive,
        maxInclusive, maxExclusive, totalDigits, fractionDigits,
        length, maxLength, minLength within datatype,
        and the min- and max- variants of Inclusive and Exclusive
        are mutually exclusive. On the other hand,  pattern and
        enumeration may repeat.
  -->
<!ENTITY % minBound "(%minInclusive; | %minExclusive;)">
<!ENTITY % maxBound "(%maxInclusive; | %maxExclusive;)">
<!ENTITY % bounds "%minBound; | %maxBound;">
<!ENTITY % numeric "%totalDigits; | %fractionDigits;">
<!ENTITY % ordered "%bounds; | %numeric;">
<!ENTITY % unordered
   "%pattern; | %enumeration; | %whiteSpace; | %length; |
   %maxLength; | %minLength;">
<!ENTITY % facet "%ordered; | %unordered;">
<!ENTITY % facetAttr
        "value CDATA #REQUIRED
        id ID #IMPLIED">
<!ENTITY % fixedAttr "fixed %boolean; #IMPLIED">
<!ENTITY % facetModel "(%annotation;)?">
<!ELEMENT %simpleType;
        ((%annotation;)?, (%restriction; | %list; | %union;))>
<!ATTLIST %simpleType;
    name      %NCName; #IMPLIED
    final     %simpleDerivationSet; #IMPLIED
    id        ID       #IMPLIED
    %simpleTypeAttrs;>
<!-- name is required at top level -->
<!ELEMENT %restriction; ((%annotation;)?,
                         (%restriction1; |
                          ((%simpleType;)?,(%facet;)*)),
                         (%attrDecls;))>
<!ATTLIST %restriction;
    base      %QName;                  #IMPLIED
    id        ID       #IMPLIED
    %restrictionAttrs;>
<!--
        base and simpleType child are mutually exclusive,
        one is required.

        restriction is shared between simpleType and
        simpleContent and complexContent (in XMLSchema.xsd).
        restriction1 is for the latter cases, when this
        is restricting a complex type, as is attrDecls.
  -->
<!ELEMENT %list; ((%annotation;)?,(%simpleType;)?)>
<!ATTLIST %list;
    itemType      %QName;             #IMPLIED
    id        ID       #IMPLIED
    %listAttrs;>
<!--
        itemType and simpleType child are mutually exclusive,
        one is required
  -->
<!ELEMENT %union; ((%annotation;)?,(%simpleType;)*)>
<!ATTLIST %union;
    id            ID       #IMPLIED
    memberTypes   %QNames;            #IMPLIED
    %unionAttrs;>
<!--
        At least one item in memberTypes or one simpleType
        child is required
  -->

<!ELEMENT %maxExclusive; %facetModel;>
<!ATTLIST %maxExclusive;
        %facetAttr;
        %fixedAttr;
        %maxExclusiveAttrs;>
<!ELEMENT %minExclusive; %facetModel;>
<!ATTLIST %minExclusive;
        %facetAttr;
        %fixedAttr;
        %minExclusiveAttrs;>

<!ELEMENT %maxInclusive; %facetModel;>
<!ATTLIST %maxInclusive;
        %facetAttr;
        %fixedAttr;
        %maxInclusiveAttrs;>
<!ELEMENT %minInclusive; %facetModel;>
<!ATTLIST %minInclusive;
        %facetAttr;
        %fixedAttr;
        %minInclusiveAttrs;>

<!ELEMENT %totalDigits; %facetModel;>
<!ATTLIST %totalDigits;
        %facetAttr;
        %fixedAttr;
        %totalDigitsAttrs;>
<!ELEMENT %fractionDigits; %facetModel;>
<!ATTLIST %fractionDigits;
        %facetAttr;
        %fixedAttr;
        %fractionDigitsAttrs;>

<!ELEMENT %length; %facetModel;>
<!ATTLIST %length;
        %facetAttr;
        %fixedAttr;
        %lengthAttrs;>
<!ELEMENT %minLength; %facetModel;>
<!ATTLIST %minLength;
        %facetAttr;
        %fixedAttr;
        %minLengthAttrs;>
<!ELEMENT %maxLength; %facetModel;>
<!ATTLIST %maxLength;
        %facetAttr;
        %fixedAttr;
        %maxLengthAttrs;>

<!-- This one can be repeated -->
<!ELEMENT %enumeration; %facetModel;>
<!ATTLIST %enumeration;
        %facetAttr;
        %enumerationAttrs;>

<!ELEMENT %whiteSpace; %facetModel;>
<!ATTLIST %whiteSpace;
        %facetAttr;
        %fixedAttr;
        %whiteSpaceAttrs;>

<!-- This one can be repeated -->
<!ELEMENT %pattern; %facetModel;>
<!ATTLIST %pattern;
        %facetAttr;
        %patternAttrs;>

## C Datatypes and Facets

### C.1 Fundamental Facets

The following table shows the values of the fundamental facets for each [·built-in·](https://www.w3.org/TR/xmlschema-2/#dt-built-in) datatype.

||Datatype|[ordered](https://www.w3.org/TR/xmlschema-2/#dt-ordered)|[bounded](https://www.w3.org/TR/xmlschema-2/#dt-bounded)|[cardinality](https://www.w3.org/TR/xmlschema-2/#dt-cardinality)|[numeric](https://www.w3.org/TR/xmlschema-2/#dt-numeric)|
|---|---|---|---|---|---|
|[primitive](https://www.w3.org/TR/xmlschema-2/#dt-primitive)|[string](https://www.w3.org/TR/xmlschema-2/#string)|false|false|countably infinite|false|
|[boolean](https://www.w3.org/TR/xmlschema-2/#boolean)|false|false|finite|false|
|[float](https://www.w3.org/TR/xmlschema-2/#float)|partial|true|finite|true|
|[double](https://www.w3.org/TR/xmlschema-2/#double)|partial|true|finite|true|
|[decimal](https://www.w3.org/TR/xmlschema-2/#decimal)|total|false|countably infinite|true|
|[duration](https://www.w3.org/TR/xmlschema-2/#duration)|partial|false|countably infinite|false|
|[dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime)|partial|false|countably infinite|false|
|[time](https://www.w3.org/TR/xmlschema-2/#time)|partial|false|countably infinite|false|
|[date](https://www.w3.org/TR/xmlschema-2/#date)|partial|false|countably infinite|false|
|[gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth)|partial|false|countably infinite|false|
|[gYear](https://www.w3.org/TR/xmlschema-2/#gYear)|partial|false|countably infinite|false|
|[gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay)|partial|false|countably infinite|false|
|[gDay](https://www.w3.org/TR/xmlschema-2/#gDay)|partial|false|countably infinite|false|
|[gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth)|partial|false|countably infinite|false|
|[hexBinary](https://www.w3.org/TR/xmlschema-2/#hexBinary)|false|false|countably infinite|false|
|[base64Binary](https://www.w3.org/TR/xmlschema-2/#base64Binary)|false|false|countably infinite|false|
|[anyURI](https://www.w3.org/TR/xmlschema-2/#anyURI)|false|false|countably infinite|false|
|[QName](https://www.w3.org/TR/xmlschema-2/#QName)|false|false|countably infinite|false|
|[NOTATION](https://www.w3.org/TR/xmlschema-2/#NOTATION)|false|false|countably infinite|false|
||   |   |   |   |   |   |
|[derived](https://www.w3.org/TR/xmlschema-2/#dt-derived)|[normalizedString](https://www.w3.org/TR/xmlschema-2/#normalizedString)|false|false|countably infinite|false|
|[token](https://www.w3.org/TR/xmlschema-2/#token)|false|false|countably infinite|false|
|[language](https://www.w3.org/TR/xmlschema-2/#language)|false|false|countably infinite|false|
|[IDREFS](https://www.w3.org/TR/xmlschema-2/#IDREFS)|false|false|countably infinite|false|
|[ENTITIES](https://www.w3.org/TR/xmlschema-2/#ENTITIES)|false|false|countably infinite|false|
|[NMTOKEN](https://www.w3.org/TR/xmlschema-2/#NMTOKEN)|false|false|countably infinite|false|
|[NMTOKENS](https://www.w3.org/TR/xmlschema-2/#NMTOKENS)|false|false|countably infinite|false|
|[Name](https://www.w3.org/TR/xmlschema-2/#Name)|false|false|countably infinite|false|
|[NCName](https://www.w3.org/TR/xmlschema-2/#NCName)|false|false|countably infinite|false|
|[ID](https://www.w3.org/TR/xmlschema-2/#ID)|false|false|countably infinite|false|
|[IDREF](https://www.w3.org/TR/xmlschema-2/#IDREF)|false|false|countably infinite|false|
|[ENTITY](https://www.w3.org/TR/xmlschema-2/#ENTITY)|false|false|countably infinite|false|
|[integer](https://www.w3.org/TR/xmlschema-2/#integer)|total|false|countably infinite|true|
|[nonPositiveInteger](https://www.w3.org/TR/xmlschema-2/#nonPositiveInteger)|total|false|countably infinite|true|
|[negativeInteger](https://www.w3.org/TR/xmlschema-2/#negativeInteger)|total|false|countably infinite|true|
|[long](https://www.w3.org/TR/xmlschema-2/#long)|total|true|finite|true|
|[int](https://www.w3.org/TR/xmlschema-2/#int)|total|true|finite|true|
|[short](https://www.w3.org/TR/xmlschema-2/#short)|total|true|finite|true|
|[byte](https://www.w3.org/TR/xmlschema-2/#byte)|total|true|finite|true|
|[nonNegativeInteger](https://www.w3.org/TR/xmlschema-2/#nonNegativeInteger)|total|false|countably infinite|true|
|[unsignedLong](https://www.w3.org/TR/xmlschema-2/#unsignedLong)|total|true|finite|true|
|[unsignedInt](https://www.w3.org/TR/xmlschema-2/#unsignedInt)|total|true|finite|true|
|[unsignedShort](https://www.w3.org/TR/xmlschema-2/#unsignedShort)|total|true|finite|true|
|[unsignedByte](https://www.w3.org/TR/xmlschema-2/#unsignedByte)|total|true|finite|true|
|[positiveInteger](https://www.w3.org/TR/xmlschema-2/#positiveInteger)|total|false|countably infinite|true|

## D ISO 8601 Date and Time Formats

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#truncatedformats)D.1 ISO 8601 Conventions

The [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) datatypes [duration](https://www.w3.org/TR/xmlschema-2/#duration), [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), [time](https://www.w3.org/TR/xmlschema-2/#time), [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay), [gDay](https://www.w3.org/TR/xmlschema-2/#gDay), [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth) and [gYear](https://www.w3.org/TR/xmlschema-2/#gYear) use lexical formats inspired by [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601). Following [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601), the lexical forms of these datatypes can include only the characters #20 through #7F. This appendix provides more detail on the ISO formats and discusses some deviations from them for the datatypes defined in this specification.

[[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601) "specifies the representation of dates in the proleptic Gregorian calendar and times and representations of periods of time". The proleptic Gregorian calendar includes dates prior to 1582 (the year it came into use as an ecclesiastical calendar). It should be pointed out that the datatypes described in this specification do not cover all the types of data covered by [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601), nor do they support all the lexical representations for those types of data.

[[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601) lexical formats are described using "pictures" in which characters are used in place of decimal digits. The allowed decimal digits are (#x30-#x39). For the primitive datatypes [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), [time](https://www.w3.org/TR/xmlschema-2/#time), [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay), [gDay](https://www.w3.org/TR/xmlschema-2/#gDay), [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth) and [gYear](https://www.w3.org/TR/xmlschema-2/#gYear). these characters have the following meanings:

- C -- represents a digit used in the thousands and hundreds components, the "century" component, of the time element "year". Legal values are from 0 to 9.
- Y -- represents a digit used in the tens and units components of the time element "year". Legal values are from 0 to 9.
- M -- represents a digit used in the time element "month". The two digits in a MM format can have values from 1 to 12.
- D -- represents a digit used in the time element "day". The two digits in a DD format can have values from 1 to 28 if the month value equals 2, 1 to 29 if the month value equals 2 and the year is a leap year, 1 to 30 if the month value equals 4, 6, 9 or 11, and 1 to 31 if the month value equals 1, 3, 5, 7, 8, 10 or 12.
- h -- represents a digit used in the time element "hour". The two digits in a hh format can have values from 0 to 24. If the value of the hour element is 24 then the values of the minutes element and the seconds element must be 00 and 00.
- m -- represents a digit used in the time element "minute". The two digits in a mm format can have values from 0 to 59.
- s -- represents a digit used in the time element "second". The two digits in a ss format can have values from 0 to 60. In the formats described in this specification the whole number of seconds [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) be followed by decimal seconds to an arbitrary level of precision. This is represented in the picture by "ss.sss". A value of 60 or more is allowed only in the case of leap seconds.

    Strictly speaking, a value of 60 or more is not sensible unless the month and day could represent March 31, June 30, September 30, or December 31 _in UTC_. Because the leap second is added or subtracted as the last second of the day in UTC time, the long (or short) minute could occur at other times in local time. In cases where the leap second is used with an inappropriate month and day it, and any fractional seconds, should considered as added or subtracted from the following minute.


For all the information items indicated by the above characters, leading zeros are required where indicated.

In addition to the above, certain characters are used as designators and appear as themselves in lexical formats.

- T -- is used as time designator to indicate the start of the representation of the time of day in [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime).
- Z -- is used as time-zone designator, immediately (without a space) following a data element expressing the time of day in Coordinated Universal Time (UTC) in [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), [time](https://www.w3.org/TR/xmlschema-2/#time), [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay), [gDay](https://www.w3.org/TR/xmlschema-2/#gDay), [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth), and [gYear](https://www.w3.org/TR/xmlschema-2/#gYear).

In the lexical format for [duration](https://www.w3.org/TR/xmlschema-2/#duration) the following characters are also used as designators and appear as themselves in lexical formats:

- P -- is used as the time duration designator, preceding a data element representing a given duration of time.
- Y -- follows the number of years in a time duration.
- M -- follows the number of months or minutes in a time duration.
- D -- follows the number of days in a time duration.
- H -- follows the number of hours in a time duration.
- S -- follows the number of seconds in a time duration.

The values of the Year, Month, Day, Hour and Minutes components are not restricted but allow an arbitrary integer. Similarly, the value of the Seconds component allows an arbitrary decimal. Thus, the lexical format for [duration](https://www.w3.org/TR/xmlschema-2/#duration) and datatypes derived from it does not follow the alternative format of § 5.5.3.2.1 of [[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#formatdetails) [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#deviantformats)D.2 Truncated and Reduced Formats

[[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601) supports a variety of "truncated" formats in which some of the characters on the left of specific formats, for example, the century, can be omitted. Truncated formats are, in general, not permitted for the datatypes defined in this specification with three exceptions. The [time](https://www.w3.org/TR/xmlschema-2/#time) datatype uses a truncated format for [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) which represents an instant of time that recurs every day. Similarly, the [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay) and [gDay](https://www.w3.org/TR/xmlschema-2/#gDay) datatypes use left-truncated formats for [date](https://www.w3.org/TR/xmlschema-2/#date). The datatype [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth) uses a right and left truncated format for [date](https://www.w3.org/TR/xmlschema-2/#date).

[[ISO 8601]](https://www.w3.org/TR/xmlschema-2/#ISO8601) also supports a variety of "reduced" or right-truncated formats in which some of the characters to the right of specific formats, such as the time specification, can be omitted. Right truncated formats are also, in general, not permitted for the datatypes defined in this specification with the following exceptions: right-truncated representations of [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) are used as lexical representations for [date](https://www.w3.org/TR/xmlschema-2/#date), [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth), [gYear](https://www.w3.org/TR/xmlschema-2/#gYear).

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#truncatedformats)D.3 Deviations from ISO 8601 Formats

        D.3.1 [Sign Allowed](https://www.w3.org/TR/xmlschema-2/#signallowed)
        D.3.2 [No Year Zero](https://www.w3.org/TR/xmlschema-2/#noYearZero)
        D.3.3 [More Than 9999 Years](https://www.w3.org/TR/xmlschema-2/#morethan9999years)
        D.3.4 [Time zone permitted](https://www.w3.org/TR/xmlschema-2/#timeZonePermited)

#### D.3.1 Sign Allowed

An optional minus sign is allowed immediately preceding, without a space, the lexical representations for [duration](https://www.w3.org/TR/xmlschema-2/#duration), [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gYear](https://www.w3.org/TR/xmlschema-2/#gYear).

#### D.3.2 No Year Zero

The year "0000" is an illegal year value.

#### D.3.3 More Than 9999 Years

To accommodate year values greater than 9999, more than four digits are allowed in the year representations of [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime), [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), and [gYear](https://www.w3.org/TR/xmlschema-2/#gYear). This follows [[ISO 8601:2000 Second Edition]](https://www.w3.org/TR/xmlschema-2/#ISO8601-2000).

#### D.3.4 Time zone permitted

The lexical representations for the datatypes [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gMonthDay](https://www.w3.org/TR/xmlschema-2/#gMonthDay), [gDay](https://www.w3.org/TR/xmlschema-2/#gDay), [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth) and [gYear](https://www.w3.org/TR/xmlschema-2/#gYear) permit an optional trailing time zone specificiation.

## E Adding durations to dateTimes

Given a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) S and a [duration](https://www.w3.org/TR/xmlschema-2/#duration) D, this appendix specifies how to compute a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) E where E is the end of the time period with start S and duration D i.e. E = S + D. Such computations are used, for example, to determine whether a [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) is within a specific time period. This appendix also addresses the addition of [duration](https://www.w3.org/TR/xmlschema-2/#duration)s to the datatypes [date](https://www.w3.org/TR/xmlschema-2/#date), [gYearMonth](https://www.w3.org/TR/xmlschema-2/#gYearMonth), [gYear](https://www.w3.org/TR/xmlschema-2/#gYear), [gDay](https://www.w3.org/TR/xmlschema-2/#gDay) and [gMonth](https://www.w3.org/TR/xmlschema-2/#gMonth), which can be viewed as a set of [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime)s. In such cases, the addition is made to the first or starting [dateTime](https://www.w3.org/TR/xmlschema-2/#dateTime) in the set.

_This is a logical explanation of the process. Actual implementations are free to optimize as long as they produce the same results._ The calculation uses the notation S[year] to represent the year field of S, S[month] to represent the month field, and so on. It also depends on the following functions:

- fQuotient(a, b) = the greatest integer less than or equal to a/b
    - fQuotient(-1,3) = -1
    - fQuotient(0,3)...fQuotient(2,3) = 0
    - fQuotient(3,3) = 1
    - fQuotient(3.123,3) = 1
- modulo(a, b) = a - fQuotient(a,b)*b
    - modulo(-1,3) = 2
    - modulo(0,3)...modulo(2,3) = 0...2
    - modulo(3,3) = 0
    - modulo(3.123,3) = 0.123
- fQuotient(a, low, high) = fQuotient(a - low, high - low)
    - fQuotient(0, 1, 13) = -1
    - fQuotient(1, 1, 13) ... fQuotient(12, 1, 13) = 0
    - fQuotient(13, 1, 13) = 1
    - fQuotient(13.123, 1, 13) = 1
- modulo(a, low, high) = modulo(a - low, high - low) + low
    - modulo(0, 1, 13) = 12
    - modulo(1, 1, 13) ... modulo(12, 1, 13) = 1...12
    - modulo(13, 1, 13) = 1
    - modulo(13.123, 1, 13) = 1.123
- maximumDayInMonthFor(yearValue, monthValue) =
    - M := modulo(monthValue, 1, 13)
    - Y := yearValue + fQuotient(monthValue, 1, 13)
    - Return a value based on M and Y:

|   |   |   |
|---|---|---|
|**31**|M = January, March, May, July, August, October, or December|   |
|**30**|M = April, June, September, or November|   |
|**29**|M = February AND (modulo(Y, 400) = 0 OR (modulo(Y, 100) != 0) AND modulo(Y, 4) = 0)|
|**28**|Otherwise|

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#adding-durations-to-instants-commutativity-associativity)E.1 Algorithm

Essentially, this calculation is equivalent to separating D into <year,month> and <day,hour,minute,second> fields. The <year,month> is added to S. If the day is out of range, it is _pinned_ to be within range. Thus April 31 turns into April 30. Then the <day,hour,minute,second> is added. This latter addition can cause the year and month to change.

Leap seconds are handled by the computation by treating them as overflows. Essentially, a value of 60 seconds in S is treated as if it were a duration of 60 seconds added to S (with a zero seconds field). All calculations thereafter use 60 seconds per minute.

Thus the addition of either PT1M or PT60S to any dateTime will always produce the same result. This is a special definition of addition which is designed to match common practice, and -- most importantly -- be stable over time.

A definition that attempted to take leap-seconds into account would need to be constantly updated, and could not predict the results of future implementation's additions. The decision to introduce a leap second in UTC is the responsibility of the [[International Earth Rotation Service (IERS)]](https://www.w3.org/TR/xmlschema-2/#IERS). They make periodic announcements as to when leap seconds are to be added, but this is not known more than a year in advance. For more information on leap seconds, see [[U.S. Naval Observatory Time Service Department]](https://www.w3.org/TR/xmlschema-2/#USNavy).

The following is the precise specification. These steps must be followed in the same order. If a field in D is not specified, it is treated as if it were zero. If a field in S is not specified, it is treated in the calculation as if it were the minimum allowed value in that field, however, after the calculation is concluded, the corresponding field in E is removed (set to unspecified).

- _Months (may be modified additionally below)_
    - temp := S[month] + D[month]
    - E[month] := modulo(temp, 1, 13)
    - carry := fQuotient(temp, 1, 13)
- _Years (may be modified additionally below)_
    - E[year] := S[year] + D[year] + carry
- _Zone_
    - E[zone] := S[zone]
- _Seconds_
    - temp := S[second] + D[second]
    - E[second] := modulo(temp, 60)
    - carry := fQuotient(temp, 60)
- _Minutes_
    - temp := S[minute] + D[minute] + carry
    - E[minute] := modulo(temp, 60)
    - carry := fQuotient(temp, 60)
- _Hours_
    - temp := S[hour] + D[hour] + carry
    - E[hour] := modulo(temp, 24)
    - carry := fQuotient(temp, 24)
- _Days_
    - if S[day] > maximumDayInMonthFor(E[year], E[month])
        - tempDays := maximumDayInMonthFor(E[year], E[month])
    - else if S[day] < 1
        - tempDays := 1
    - else
        - tempDays := S[day]
    - E[day] := tempDays + D[day] + carry
    - **START LOOP**
        - **IF** E[day] < 1
            - E[day] := E[day] + maximumDayInMonthFor(E[year], E[month] - 1)
            - carry := -1
        - **ELSE IF** E[day] > maximumDayInMonthFor(E[year], E[month])
            - E[day] := E[day] - maximumDayInMonthFor(E[year], E[month])
            - carry := 1
        - **ELSE EXIT LOOP**
        - temp := E[month] + carry
        - E[month] := modulo(temp, 1, 13)
        - E[year] := E[year] + fQuotient(temp, 1, 13)
        - **GOTO START LOOP**

_Examples:_

|dateTime|duration|result|
|:-:|:-:|:-:|
|2000-01-12T12:13:14Z|P1Y3M5DT7H10M3.3S|2001-04-17T19:23:17.3Z|
|2000-01|-P3M|1999-10|
|2000-01-12|PT33H|2000-01-13|

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#)E.2 Commutativity and Associativity

Time durations are added by simply adding each of their fields, respectively, without overflow.

The order of addition of durations to instants _is_ significant. For example, there are cases where:

((dateTime + duration1) + duration2) != ((dateTime + duration2) + duration1)

_Example:_

(2000-03-30 + P1D) + P1M = 2000-03-31 + P1M = 2000-**04-30**

(2000-03-30 + P1M) + P1D = 2000-04-30 + P1D = 2000-**05-01**

## F Regular Expressions

A [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex) _R_ is a sequence of characters that denote a **set of strings**  _L(R)_. When used to constrain a [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space), a **regular expression**  _R_ asserts that only strings in _L(R)_ are valid literals for values of that type.

**Note:**  Unlike some popular regular expression languages (including those defined by Perl and standard Unix utilities), the regular expression language defined here implicitly anchors all regular expressions at the head and tail, as the most common use of regular expressions in [·pattern·](https://www.w3.org/TR/xmlschema-2/#dt-pattern) is to match entire literals. For example, a datatype [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from [string](https://www.w3.org/TR/xmlschema-2/#string) such that all values must begin with the character `A` (#x41) and end with the character `Z` (#x5a) would be defined as follows:

<simpleType name='myString'>
 <restriction base='string'>
  <pattern value='A.*Z'/>
 </restriction>
</simpleType>

In regular expression languages that are not implicitly anchored at the head and tail, it is customary to write the equivalent regular expression as:

   `^A.*Z$`

where "^" anchors the pattern at the head and "$" anchors at the tail.

In those rare cases where an unanchored match is desired, including `.*` at the beginning and ending of the regular expression will achieve the desired results. For example, a datatype [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) from string such that all values must contain at least 3 consecutive `A` (`#x41`) characters somewhere within the value could be defined as follows:

<simpleType name='myString'>
 <restriction base='string'>
  <pattern value='.*AAA.*'/>
 </restriction>
</simpleType>

[Definition:]  A **regular expression** is composed from zero or more [·branch·](https://www.w3.org/TR/xmlschema-2/#dt-branch)es, separated by `|` characters.

|   |
|---|
|**Regular Expression**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[1]\|`regExp`\|::=\|`[branch](https://www.w3.org/TR/xmlschema-2/#nt-branch) ( '\\|' [branch](https://www.w3.org/TR/xmlschema-2/#nt-branch) )*`\||

|For all [·branch·](https://www.w3.org/TR/xmlschema-2/#dt-branch)es _S_, and for all [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex)s _T_, valid [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex)s _R_ are:|Denoting the set of strings _L(R)_ containing:|
|---|---|
|(empty string)|the set containing just the empty string|
|_S_|all strings in _L(S)_|
|_S_\|_T_|all strings in _L(S)_ and all strings in _L(T)_|

[Definition:]   A **branch** consists of zero or more [·piece·](https://www.w3.org/TR/xmlschema-2/#dt-piece)s, concatenated together.

|   |
|---|
|**Branch**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[2]\|`branch`\|::=\|`[piece](https://www.w3.org/TR/xmlschema-2/#nt-piece)*`\||

|For all [·piece·](https://www.w3.org/TR/xmlschema-2/#dt-piece)s _S_, and for all [·branch·](https://www.w3.org/TR/xmlschema-2/#dt-branch)es _T_, valid [·branch·](https://www.w3.org/TR/xmlschema-2/#dt-branch)es _R_ are:|Denoting the set of strings _L(R)_ containing:|
|---|---|
|_S_|all strings in _L(S)_|
|_S__T_|all strings _st_ with _s_ in _L(S)_ and _t_ in _L(T)_|

[Definition:]   A **piece** is an [·atom·](https://www.w3.org/TR/xmlschema-2/#dt-atom), possibly followed by a [·quantifier·](https://www.w3.org/TR/xmlschema-2/#dt-quantifier).

|   |
|---|
|**Piece**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[3]\|`piece`\|::=\|`[atom](https://www.w3.org/TR/xmlschema-2/#nt-atom) [quantifier](https://www.w3.org/TR/xmlschema-2/#nt-quantifier)?`\||

|For all [·atom·](https://www.w3.org/TR/xmlschema-2/#dt-atom)s _S_ and non-negative integers _n_, _m_ such that _n <= m_, valid [·piece·](https://www.w3.org/TR/xmlschema-2/#dt-piece)s _R_ are:|Denoting the set of strings _L(R)_ containing:|
|---|---|
|_S_|all strings in _L(S)_|
|_S_?|the empty string, and all strings in _L(S)_.|
|_S_*|All strings in _L(S?)_ and all strings _st_ with _s_ in _L(S*)_ and _t_ in _L(S)_. _( all concatenations of zero or more strings from L(S) )_|
|_S_+|All strings _st_ with _s_ in _L(S)_ and _t_ in _L(S*)_. _( all concatenations of one or more strings from L(S) )_|
|_S_{n,m}|All strings _st_ with _s_ in _L(S)_ and _t_ in _L(S{n-1,m-1})_. _( All sequences of at least n, and at most m, strings from L(S) )_|
|_S_{n}|All strings in _L(S{n,n})_. _( All sequences of exactly n strings from L(S) )_|
|_S_{n,}|All strings in L(S{n}S*) _( All sequences of at least n, strings from L(S) )_|
|_S_{0,m}|All strings _st_ with _s_ in _L(S?)_ and _t_ in _L(S{0,m-1})_. _( All sequences of at most m, strings from L(S) )_|
|_S_{0,0}|The set containing only the empty string|

**Note:**  The regular expression language in the Perl Programming Language [[Perl]](https://www.w3.org/TR/xmlschema-2/#Perl) does not include a quantifier of the form `S{,m}`, since it is logically equivalent to `S{0,m}`. We have, therefore, left this logical possibility out of the regular expression language defined by this specification.

[Definition:]   A **quantifier** is one of `?`, `*`, `+`, `{n,m}` or `{n,}`, which have the meanings defined in the table above.

|   |
|---|
|**Quanitifer**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[4]\|`quantifier`\|::=\|`[?*+] \\| ( '{' [quantity](https://www.w3.org/TR/xmlschema-2/#nt-quantity) '}' )`\|<br>\|[5]\|`quantity`\|::=\|`[quantRange](https://www.w3.org/TR/xmlschema-2/#nt-quantRange) \\| [quantMin](https://www.w3.org/TR/xmlschema-2/#nt-quantMin) \\| [QuantExact](https://www.w3.org/TR/xmlschema-2/#nt-QuantExact)`\|<br>\|[6]\|`quantRange`\|::=\|`[QuantExact](https://www.w3.org/TR/xmlschema-2/#nt-QuantExact) ',' [QuantExact](https://www.w3.org/TR/xmlschema-2/#nt-QuantExact)`\|<br>\|[7]\|`quantMin`\|::=\|`[QuantExact](https://www.w3.org/TR/xmlschema-2/#nt-QuantExact) ','`\|<br>\|[8]\|`QuantExact`\|::=\|`[0-9]+`\||

[Definition:]   An **atom** is either a [·normal character·](https://www.w3.org/TR/xmlschema-2/#dt-normalc), a [·character class·](https://www.w3.org/TR/xmlschema-2/#dt-charclass), or a parenthesized [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

|   |
|---|
|**Atom**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[9]\|`atom`\|::=\|`[Char](https://www.w3.org/TR/xmlschema-2/#nt-Char) \\| [charClass](https://www.w3.org/TR/xmlschema-2/#nt-charClass) \\| ( '(' [regExp](https://www.w3.org/TR/xmlschema-2/#nt-regExp) ')' )`\||

|For all [·normal character·](https://www.w3.org/TR/xmlschema-2/#dt-normalc)s _c_, [·character class·](https://www.w3.org/TR/xmlschema-2/#dt-charclass)es _C_, and [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex)s _S_, valid [·atom·](https://www.w3.org/TR/xmlschema-2/#dt-atom)s _R_ are:|Denoting the set of strings _L(R)_ containing:|
|---|---|
|_c_|the single string consisting only of _c_|
|_C_|all strings in _L(C)_|
|(_S_)|all strings in _L(S)_|

[Definition:]   A **metacharacter** is either `.`, `\`, `?`, `*`, `+`, `{`, `}` `(`, `)`, `[` or `]`. These characters have special meanings in [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex)s, but can be escaped to form [·atom·](https://www.w3.org/TR/xmlschema-2/#dt-atom)s that denote the sets of strings containing only themselves, i.e., an escaped [·metacharacter·](https://www.w3.org/TR/xmlschema-2/#dt-metac) behaves like a [·normal character·](https://www.w3.org/TR/xmlschema-2/#dt-normalc).

[Definition:]   A **normal character** is any XML character that is not a metacharacter. In [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex)s, a normal character is an atom that denotes the singleton set of strings containing only itself.

|   |
|---|
|**Normal Character**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[10]\|`Char`\|::=\|`[^.\?*+()\\|#x5B#x5D]`\||

Note that a [·normal character·](https://www.w3.org/TR/xmlschema-2/#dt-normalc) can be represented either as itself, or with a [character reference](https://www.w3.org/TR/2000/WD-xml-2e-20000814#dt-charref).

### F.1 Character Classes

[Definition:]   A **character class** is an [·atom·](https://www.w3.org/TR/xmlschema-2/#dt-atom)  _R_ that identifies a **set of characters**  _C(R)_. The set of strings _L(R)_ denoted by a character class _R_ contains one single-character string "_c_" for each character _c_ in _C(R)_.

|   |
|---|
|**Character Class**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[11]\|`charClass`\|::=\|`[charClassEsc](https://www.w3.org/TR/xmlschema-2/#nt-charClassEsc) \\| [charClassExpr](https://www.w3.org/TR/xmlschema-2/#nt-charClassExpr) \\| [WildcardEsc](https://www.w3.org/TR/xmlschema-2/#nt-WildcardEsc)`\||

A character class is either a [·character class escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces) or a [·character class expression·](https://www.w3.org/TR/xmlschema-2/#dt-charexpr).

[Definition:]   A **character class expression** is a [·character group·](https://www.w3.org/TR/xmlschema-2/#dt-chargroup) surrounded by `[` and `]` characters. For all character groups _G_, [_G_] is a valid **character class expression**, identifying the set of characters _C_([_G_]) = _C_(_G_).

|   |
|---|
|**Character Class Expression**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[12]\|`charClassExpr`\|::=\|`'[' [charGroup](https://www.w3.org/TR/xmlschema-2/#nt-charGroup) ']'`\||

[Definition:]   A **character group** is either a [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup), a [·negative character group·](https://www.w3.org/TR/xmlschema-2/#dt-negchargroup), or a [·character class subtraction·](https://www.w3.org/TR/xmlschema-2/#dt-subchargroup).

|   |
|---|
|**Character Group**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[13]\|`charGroup`\|::=\|`[posCharGroup](https://www.w3.org/TR/xmlschema-2/#nt-posCharGroup) \\| [negCharGroup](https://www.w3.org/TR/xmlschema-2/#nt-negCharGroup) \\| [charClassSub](https://www.w3.org/TR/xmlschema-2/#nt-charClassSub)`\||

[Definition:]   A **positive character group** consists of one or more [·character range·](https://www.w3.org/TR/xmlschema-2/#dt-charrange)s or [·character class escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces)s, concatenated together. A **positive character group** identifies the set of characters containing all of the characters in all of the sets identified by its constituent ranges or escapes.

|   |
|---|
|**Positive Character Group**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[14]\|`posCharGroup`\|::=\|`( [charRange](https://www.w3.org/TR/xmlschema-2/#nt-charRange) \\| [charClassEsc](https://www.w3.org/TR/xmlschema-2/#nt-charClassEsc) )+`\||

|For all [·character range·](https://www.w3.org/TR/xmlschema-2/#dt-charrange)s _R_, all [·character class escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces)s _E_, and all [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup)s _P_, valid [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup)s _G_ are:|Identifying the set of characters _C(G)_ containing:|
|---|---|
|_R_|all characters in _C(R)_.|
|_E_|all characters in _C(E)_.|
|_RP_|all characters in _C(R)_ and all characters in _C(P)_.|
|_EP_|all characters in _C(E)_ and all characters in _C(P)_.|

[Definition:]   A **negative character group** is a [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup) preceded by the `^` character. For all [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup)s _P_, ^_P_ is a valid **negative character group**, and _C(^P)_ contains all XML characters that are _not_ in _C(P)_.

|   |
|---|
|**Negative Character Group**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[15]\|`negCharGroup`\|::=\|`'^' [posCharGroup](https://www.w3.org/TR/xmlschema-2/#nt-posCharGroup)`\||

[Definition:]   A **character class subtraction** is a [·character class expression·](https://www.w3.org/TR/xmlschema-2/#dt-charexpr) subtracted from a [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup) or [·negative character group·](https://www.w3.org/TR/xmlschema-2/#dt-negchargroup), using the `-` character.

|   |
|---|
|**Character Class Subtraction**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[16]\|`charClassSub`\|::=\|`( [posCharGroup](https://www.w3.org/TR/xmlschema-2/#nt-posCharGroup) \\| [negCharGroup](https://www.w3.org/TR/xmlschema-2/#nt-negCharGroup) ) '-' [charClassExpr](https://www.w3.org/TR/xmlschema-2/#nt-charClassExpr)`\||

For any [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup) or [·negative character group·](https://www.w3.org/TR/xmlschema-2/#dt-negchargroup) _G_, and any [·character class expression·](https://www.w3.org/TR/xmlschema-2/#dt-charexpr) _C_, _G-C_ is a valid [·character class subtraction·](https://www.w3.org/TR/xmlschema-2/#dt-subchargroup), identifying the set of all characters in _C(G)_ that are not also in _C(C)_.

[Definition:]   A **character range** _R_ identifies a set of characters _C(R)_ containing all XML characters with UCS code points in a specified range.

|   |
|---|
|**Character Range**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[17]\|`charRange`\|::=\|`[seRange](https://www.w3.org/TR/xmlschema-2/#nt-seRange) \\| [XmlCharIncDash](https://www.w3.org/TR/xmlschema-2/#nt-XmlCharIncDash)`\|<br>\|[18]\|`seRange`\|::=\|`[charOrEsc](https://www.w3.org/TR/xmlschema-2/#nt-charOrEsc) '-' [charOrEsc](https://www.w3.org/TR/xmlschema-2/#nt-charOrEsc)`\|<br>\|[20]\|`charOrEsc`\|::=\|`[XmlChar](https://www.w3.org/TR/xmlschema-2/#nt-XmlChar) \\| [SingleCharEsc](https://www.w3.org/TR/xmlschema-2/#nt-SingleCharEsc)`\|<br>\|[21]\|`XmlChar`\|::=\|`[^\#x2D#x5B#x5D]`\|<br>\|[22]\|`XmlCharIncDash`\|::=\|`[^\#x5B#x5D]`\||

A single XML character is a [·character range·](https://www.w3.org/TR/xmlschema-2/#dt-charrange) that identifies the set of characters containing only itself. All XML characters are valid character ranges, except as follows:

- The `[`, `]`, `-` and `\` characters are not valid character ranges;
- The `^` character is only valid at the beginning of a [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup) if it is part of a [·negative character group·](https://www.w3.org/TR/xmlschema-2/#dt-negchargroup)
- The `-` character is a valid character range only at the beginning or end of a [·positive character group·](https://www.w3.org/TR/xmlschema-2/#dt-poschargroup).

**Note:** The grammar for [·character range·](https://www.w3.org/TR/xmlschema-2/#dt-charrange) as given above is ambiguous, but the second and third bullets above together remove the ambiguity.

A [·character range·](https://www.w3.org/TR/xmlschema-2/#dt-charrange) [·may·](https://www.w3.org/TR/xmlschema-2/#dt-may) also be written in the form _s-e_, identifying the set that contains all XML characters with UCS code points greater than or equal to the code point of _s_, but not greater than the code point of _e_.

_s-e_ is a valid character range iff:

- _s_ is a [·single character escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces1), or an XML character;
- _s_ is not `\`
- If s is the first character in a [·character class expression·](https://www.w3.org/TR/xmlschema-2/#dt-charexpr), then _s_ is not `^`
- _e_ is a [·single character escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces1), or an XML character;
- _e_ is not `\` or `[`; and
- The code point of _e_ is greater than or equal to the code point of _s_;

**Note:**  The code point of a [·single character escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces1) is the code point of the single character in the set of characters that it identifies.

#### F.1.1 Character Class Escapes

[Definition:]   A **character class escape** is a short sequence of characters that identifies predefined character class. The valid character class escapes are the [·single character escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces1)s, the [·multi-character escape·](https://www.w3.org/TR/xmlschema-2/#dt-ccesN)s, and the [·category escape·](https://www.w3.org/TR/xmlschema-2/#dt-ccescat)s (including the [·block escape·](https://www.w3.org/TR/xmlschema-2/#dt-ccesblock)s).

|   |
|---|
|**Character Class Escape**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[23]\|`charClassEsc`\|::=\|`( [SingleCharEsc](https://www.w3.org/TR/xmlschema-2/#nt-SingleCharEsc) \\| [MultiCharEsc](https://www.w3.org/TR/xmlschema-2/#nt-MultiCharEsc) \\| [catEsc](https://www.w3.org/TR/xmlschema-2/#nt-catEsc) \\| [complEsc](https://www.w3.org/TR/xmlschema-2/#nt-complEsc) )`\||

[Definition:]   A **single character escape** identifies a set containing a only one character -- usually because that character is difficult or impossible to write directly into a [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex).

|   |
|---|
|**Single Character Escape**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[24]\|`SingleCharEsc`\|::=\|`'\' [nrt\\\|.?*+(){}#x2D#x5B#x5D#x5E]`\||

|The valid [·single character escape·](https://www.w3.org/TR/xmlschema-2/#dt-cces1)s are:|Identifying the set of characters _C(R)_ containing:|
|---|---|
|`\n`|the newline character (#xA)|
|`\r`|the return character (#xD)|
|`\t`|the tab character (#x9)|
|`\\`|\|
|`\\|`|\||
|`\.`|.|
|`\-`|-|
|`\^`|^|
|`\?`|?|
|`\*`|*|
|`\+`|+|
|`\{`|{|
|`\}`|}|
|`\(`|(|
|`\)`|)|
|`\[`|[|
|`\]`|]|

[Definition:]   [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB) specifies a number of possible values for the "General Category" property and provides mappings from code points to specific character properties. The set containing all characters that have property `X`, can be identified with a **category escape** `\p{X}`. The complement of this set is specified with the **category escape** `\P{X}`. (`[\P{X}]` = `[^\p{X}]`).

|   |
|---|
|**Category Escape**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[25]\|`catEsc`\|::=\|`'\p{' [charProp](https://www.w3.org/TR/xmlschema-2/#nt-charProp) '}'`\|<br>\|[26]\|`complEsc`\|::=\|`'\P{' [charProp](https://www.w3.org/TR/xmlschema-2/#nt-charProp) '}'`\|<br>\|[27]\|`charProp`\|::=\|`[IsCategory](https://www.w3.org/TR/xmlschema-2/#nt-IsCategory) \\| [IsBlock](https://www.w3.org/TR/xmlschema-2/#nt-IsBlock)`\||

**Note:**  [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB) is subject to future revision. For example, the mapping from code points to character properties might be updated. All [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) support the character properties defined in the version of [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB) that is current at the time this specification became a W3C Recommendation. However, implementors are encouraged to support the character properties defined in any future version.

The following table specifies the recognized values of the "General Category" property.

|Category|Property|Meaning|
|---|---|---|
|Letters|L|All Letters|
|Lu|uppercase|
|Ll|lowercase|
|Lt|titlecase|
|Lm|modifier|
|Lo|other|
||   |   |
|Marks|M|All Marks|
|Mn|nonspacing|
|Mc|spacing combining|
|Me|enclosing|
||   |   |
|Numbers|N|All Numbers|
|Nd|decimal digit|
|Nl|letter|
|No|other|
||   |   |
|Punctuation|P|All Punctuation|
|Pc|connector|
|Pd|dash|
|Ps|open|
|Pe|close|
|Pi|initial quote (may behave like Ps or Pe depending on usage)|
|Pf|final quote (may behave like Ps or Pe depending on usage)|
|Po|other|
||   |   |
|Separators|Z|All Separators|
|Zs|space|
|Zl|line|
|Zp|paragraph|
||   |   |
|Symbols|S|All Symbols|
|Sm|math|
|Sc|currency|
|Sk|modifier|
|So|other|
||   |   |
|Other|C|All Others|
|Cc|control|
|Cf|format|
|Co|private use|
|Cn|not assigned|

|   |
|---|
|**Categories**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[28]\|`IsCategory`\|::=\|`[Letters](https://www.w3.org/TR/xmlschema-2/#nt-Letters) \\| [Marks](https://www.w3.org/TR/xmlschema-2/#nt-Marks) \\| [Numbers](https://www.w3.org/TR/xmlschema-2/#nt-Numbers) \\| [Punctuation](https://www.w3.org/TR/xmlschema-2/#nt-Punctuation) \\| [Separators](https://www.w3.org/TR/xmlschema-2/#nt-Separators) \\| [Symbols](https://www.w3.org/TR/xmlschema-2/#nt-Symbols) \\| [Others](https://www.w3.org/TR/xmlschema-2/#nt-Others)`\|<br>\|[29]\|`Letters`\|::=\|`'L' [ultmo]?`\|<br>\|[30]\|`Marks`\|::=\|`'M' [nce]?`\|<br>\|[31]\|`Numbers`\|::=\|`'N' [dlo]?`\|<br>\|[32]\|`Punctuation`\|::=\|`'P' [cdseifo]?`\|<br>\|[33]\|`Separators`\|::=\|`'Z' [slp]?`\|<br>\|[34]\|`Symbols`\|::=\|`'S' [mcko]?`\|<br>\|[35]\|`Others`\|::=\|`'C' [cfon]?`\||

**Note:**  The properties mentioned above exclude the `Cs` property. The `Cs` property identifies "surrogate" characters, which do not occur at the level of the "character abstraction" that XML instance documents operate on.

[Definition:]   [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB) groups code points into a number of blocks such as Basic Latin (i.e., ASCII), Latin-1 Supplement, Hangul Jamo, CJK Compatibility, etc. The set containing all characters that have block name `X` (with all white space stripped out), can be identified with a **block escape** `\p{IsX}`. The complement of this set is specified with the **block escape** `\P{IsX}`. (`[\P{IsX}]` = `[^\p{IsX}]`).

|   |
|---|
|**Block Escape**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[36]\|`IsBlock`\|::=\|`'Is' [a-zA-Z0-9#x2D]+`\||

The following table specifies the recognized block names (for more information, see the "Blocks.txt" file in [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB)).

|Start Code|End Code|Block Name||Start Code|End Code|Block Name|
|---|---|---|---|---|---|---|
|#x0000|#x007F|BasicLatin||#x0080|#x00FF|Latin-1Supplement|
|#x0100|#x017F|LatinExtended-A||#x0180|#x024F|LatinExtended-B|
|#x0250|#x02AF|IPAExtensions||#x02B0|#x02FF|SpacingModifierLetters|
|#x0300|#x036F|CombiningDiacriticalMarks||#x0370|#x03FF|Greek|
|#x0400|#x04FF|Cyrillic||#x0530|#x058F|Armenian|
|#x0590|#x05FF|Hebrew||#x0600|#x06FF|Arabic|
|#x0700|#x074F|Syriac||#x0780|#x07BF|Thaana|
|#x0900|#x097F|Devanagari||#x0980|#x09FF|Bengali|
|#x0A00|#x0A7F|Gurmukhi||#x0A80|#x0AFF|Gujarati|
|#x0B00|#x0B7F|Oriya||#x0B80|#x0BFF|Tamil|
|#x0C00|#x0C7F|Telugu||#x0C80|#x0CFF|Kannada|
|#x0D00|#x0D7F|Malayalam||#x0D80|#x0DFF|Sinhala|
|#x0E00|#x0E7F|Thai||#x0E80|#x0EFF|Lao|
|#x0F00|#x0FFF|Tibetan||#x1000|#x109F|Myanmar|
|#x10A0|#x10FF|Georgian||#x1100|#x11FF|HangulJamo|
|#x1200|#x137F|Ethiopic||#x13A0|#x13FF|Cherokee|
|#x1400|#x167F|UnifiedCanadianAboriginalSyllabics||#x1680|#x169F|Ogham|
|#x16A0|#x16FF|Runic||#x1780|#x17FF|Khmer|
|#x1800|#x18AF|Mongolian||#x1E00|#x1EFF|LatinExtendedAdditional|
|#x1F00|#x1FFF|GreekExtended||#x2000|#x206F|GeneralPunctuation|
|#x2070|#x209F|SuperscriptsandSubscripts||#x20A0|#x20CF|CurrencySymbols|
|#x20D0|#x20FF|CombiningMarksforSymbols||#x2100|#x214F|LetterlikeSymbols|
|#x2150|#x218F|NumberForms||#x2190|#x21FF|Arrows|
|#x2200|#x22FF|MathematicalOperators||#x2300|#x23FF|MiscellaneousTechnical|
|#x2400|#x243F|ControlPictures||#x2440|#x245F|OpticalCharacterRecognition|
|#x2460|#x24FF|EnclosedAlphanumerics||#x2500|#x257F|BoxDrawing|
|#x2580|#x259F|BlockElements||#x25A0|#x25FF|GeometricShapes|
|#x2600|#x26FF|MiscellaneousSymbols||#x2700|#x27BF|Dingbats|
|#x2800|#x28FF|BraillePatterns||#x2E80|#x2EFF|CJKRadicalsSupplement|
|#x2F00|#x2FDF|KangxiRadicals||#x2FF0|#x2FFF|IdeographicDescriptionCharacters|
|#x3000|#x303F|CJKSymbolsandPunctuation||#x3040|#x309F|Hiragana|
|#x30A0|#x30FF|Katakana||#x3100|#x312F|Bopomofo|
|#x3130|#x318F|HangulCompatibilityJamo||#x3190|#x319F|Kanbun|
|#x31A0|#x31BF|BopomofoExtended||#x3200|#x32FF|EnclosedCJKLettersandMonths|
|#x3300|#x33FF|CJKCompatibility||#x3400|#x4DB5|CJKUnifiedIdeographsExtensionA|
|#x4E00|#x9FFF|CJKUnifiedIdeographs||#xA000|#xA48F|YiSyllables|
|#xA490|#xA4CF|YiRadicals||#xAC00|#xD7A3|HangulSyllables|
||||||||
|||||#xE000|#xF8FF|PrivateUse|
|#xF900|#xFAFF|CJKCompatibilityIdeographs||#xFB00|#xFB4F|AlphabeticPresentationForms|
|#xFB50|#xFDFF|ArabicPresentationForms-A||#xFE20|#xFE2F|CombiningHalfMarks|
|#xFE30|#xFE4F|CJKCompatibilityForms||#xFE50|#xFE6F|SmallFormVariants|
|#xFE70|#xFEFE|ArabicPresentationForms-B||#xFEFF|#xFEFF|Specials|
|#xFF00|#xFFEF|HalfwidthandFullwidthForms||#xFFF0|#xFFFD|Specials|

**Note:**  The blocks mentioned above exclude the `HighSurrogates`, `LowSurrogates` and `HighPrivateUseSurrogates` blocks. These blocks identify "surrogate" characters, which do not occur at the level of the "character abstraction" that XML instance documents operate on.

**Note:**  [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB) is subject to future revision. For example, the grouping of code points into blocks might be updated. All [·minimally conforming·](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming) processors [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) support the blocks defined in the version of [[Unicode Database]](https://www.w3.org/TR/xmlschema-2/#UnicodeDB) that is current at the time this specification became a W3C Recommendation. However, implementors are encouraged to support the blocks defined in any future version of the Unicode Standard.

For example, the [·block escape·](https://www.w3.org/TR/xmlschema-2/#dt-ccesblock) for identifying the ASCII characters is `\p{IsBasicLatin}`.

[Definition:]   A **multi-character escape** provides a simple way to identify a commonly used set of characters:

|   |
|---|
|**Multi-Character Escape**|
|\|   \|   \|   \|   \|<br>\|---\|---\|---\|---\|<br>\|[37]\|`MultiCharEsc`\|::=\|`'\' [sSiIcCdDwW]`\|<br>\|[37a]\|`WildcardEsc`\|::=\|`'.'`\||

|Character sequence|Equivalent [·character class·](https://www.w3.org/TR/xmlschema-2/#dt-charclass)|
|---|---|
|.|[^\n\r]|
|\s|[#x20\t\n\r]|
|\S|[^\s]|
|\i|the set of initial name characters, those [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match)ed by [Letter](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-Letter) \| '_' \| ':'|
|\I|[^\i]|
|\c|the set of name characters, those [·match·](https://www.w3.org/TR/xmlschema-2/#dt-match)ed by [NameChar](https://www.w3.org/TR/2000/WD-xml-2e-20000814#NT-NameChar)|
|\C|[^\c]|
|\d|\p{Nd}|
|\D|[^\d]|
|\w|[#x0000-#x10FFFF]-[\p{P}\p{Z}\p{C}] (_all characters except the set of "punctuation", "separator" and "other" characters_)|
|\W|[^\w]|

**Note:**  The [·regular expression·](https://www.w3.org/TR/xmlschema-2/#dt-regex) language defined here does not attempt to provide a general solution to "regular expressions" over UCS character sequences. In particular, it does not easily provide for matching sequences of base characters and combining marks. The language is targeted at support of "Level 1" features as defined in [[Unicode Regular Expression Guidelines]](https://www.w3.org/TR/xmlschema-2/#unicodeRegEx). It is hoped that future versions of this specification will provide support for "Level 2" features.

## G Glossary (non-normative)

The listing below is for the benefit of readers of a printed version of this document: it collects together all the definitions which appear in the document above.

[atomic](https://www.w3.org/TR/xmlschema-2/#dt-atomic)

**Atomic** datatypes are those having values which are regarded by this specification as being indivisible.

[base type](https://www.w3.org/TR/xmlschema-2/#dt-basetype)

Every datatype that is [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by **restriction** is defined in terms of an existing datatype, referred to as its **base type**. **base type**s can be either [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) or [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived).

[bounded](https://www.w3.org/TR/xmlschema-2/#dt-bounded)

A datatype is **bounded** if its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) has either an [·inclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-upper-bound) or an [·exclusive upper bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-upper-bound) and either an [·inclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-lower-bound) or an [·exclusive lower bound·](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-lower-bound).

[built-in](https://www.w3.org/TR/xmlschema-2/#dt-built-in)

**Built-in** datatypes are those which are defined in this specification, and can be either [·primitive·](https://www.w3.org/TR/xmlschema-2/#dt-primitive) or [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived);

[canonical lexical representation](https://www.w3.org/TR/xmlschema-2/#dt-canonical-representation)

A **canonical lexical representation** is a set of literals from among the valid set of literals for a datatype such that there is a one-to-one mapping between literals in the **canonical lexical representation** and values in the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[cardinality](https://www.w3.org/TR/xmlschema-2/#dt-cardinality)

Every [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) has associated with it the concept of **cardinality**. Some [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s are finite, some are countably infinite while still others could conceivably be uncountably infinite (although no [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) defined by this specification is uncountable infinite). A datatype is said to have the cardinality of its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[comparable](https://www.w3.org/TR/xmlschema-2/#dt-comparable)

otherwise they are **comparable**.

[conformance to the XML Representation of Schemas](https://www.w3.org/TR/xmlschema-2/#dt-interchange)

Processors which accept schemas in the form of XML documents as described in [XML Representation of Simple Type Definition Schema Components (§4.1.2)](https://www.w3.org/TR/xmlschema-2/#xr-defn) (and other relevant portions of [Datatype components (§4)](https://www.w3.org/TR/xmlschema-2/#datatype-components)) are additionally said to provide **conformance to the XML Representation of Schemas**, and [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must), when processing schema documents, completely and correctly implement all [·Schema Representation Constraint·](https://www.w3.org/TR/xmlschema-2/#dt-src)s in this specification, and [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) adhere exactly to the specifications in [XML Representation of Simple Type Definition Schema Components (§4.1.2)](https://www.w3.org/TR/xmlschema-2/#xr-defn) (and other relevant portions of [Datatype components (§4)](https://www.w3.org/TR/xmlschema-2/#datatype-components)) for mapping the contents of such documents to [schema components](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html#key-component) for use in validation.

[constraining facet](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)

A **constraining facet** is an optional property that can be applied to a datatype to constrain its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[Constraint on Schemas](https://www.w3.org/TR/xmlschema-2/#dt-cos)

**Constraint on Schemas**

[datatype](https://www.w3.org/TR/xmlschema-2/#dt-datatype)

In this specification, a **datatype** is a 3-tuple, consisting of a) a set of distinct values, called its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), b) a set of lexical representations, called its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space), and c) a set of [·facet·](https://www.w3.org/TR/xmlschema-2/#dt-facet)s that characterize properties of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), individual values or lexical items.

[derived](https://www.w3.org/TR/xmlschema-2/#dt-derived)

**Derived** datatypes are those that are defined in terms of other datatypes.

[error](https://www.w3.org/TR/xmlschema-2/#dt-error)

**error**

[exclusive lower bound](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-lower-bound)

A value _l_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _L_ is said to be an **exclusive lower bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _L_) if for all _v_ in _V_, _l_ < _v_.

[exclusive upper bound](https://www.w3.org/TR/xmlschema-2/#dt-exclusive-upper-bound)

A value _u_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _U_ is said to be an **exclusive upper bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _U_) if for all _v_ in _V_, _u_ > _v_.

[facet](https://www.w3.org/TR/xmlschema-2/#dt-facet)

A **facet** is a single defining aspect of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space). Generally speaking, each facet characterizes a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) along independent axes or dimensions.

[for compatibility](https://www.w3.org/TR/xmlschema-2/#dt-compatibility)

for compatibility

[fundamental facet](https://www.w3.org/TR/xmlschema-2/#dt-fundamental-facet)

A **fundamental facet** is an abstract property which serves to semantically characterize the values in a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[inclusive lower bound](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-lower-bound)

A value _l_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _L_ is said to be an **inclusive lower bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _L_) if for all _v_ in _V_, _l_ <= _v_.

[inclusive upper bound](https://www.w3.org/TR/xmlschema-2/#dt-inclusive-upper-bound)

A value _u_ in an [·ordered·](https://www.w3.org/TR/xmlschema-2/#dt-ordered)  [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _U_ is said to be an **inclusive upper bound** of a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) _V_ (where _V_ is a subset of _U_) if for all _v_ in _V_, _u_ >= _v_.

[incomparable](https://www.w3.org/TR/xmlschema-2/#dt-incomparable)

When _a <> b_, _a_ and _b_ are **incomparable**,

[itemType](https://www.w3.org/TR/xmlschema-2/#dt-itemType)

The [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) or [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype that participates in the definition of a [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype is known as the **itemType** of that [·list·](https://www.w3.org/TR/xmlschema-2/#dt-list) datatype.

[lexical space](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)

A **lexical space** is the set of valid _literals_ for a datatype.

[list](https://www.w3.org/TR/xmlschema-2/#dt-list)

**List** datatypes are those having values each of which consists of a finite-length (possibly empty) sequence of values of an [·atomic·](https://www.w3.org/TR/xmlschema-2/#dt-atomic) datatype.

[match](https://www.w3.org/TR/xmlschema-2/#dt-match)

**match**

[may](https://www.w3.org/TR/xmlschema-2/#dt-may)

**may**

[memberTypes](https://www.w3.org/TR/xmlschema-2/#dt-memberTypes)

The datatypes that participate in the definition of a [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype are known as the **memberTypes** of that [·union·](https://www.w3.org/TR/xmlschema-2/#dt-union) datatype.

[minimally conforming](https://www.w3.org/TR/xmlschema-2/#dt-minimally-conforming)

**Minimally conforming** processors [·must·](https://www.w3.org/TR/xmlschema-2/#dt-must) completely and correctly implement the [·Constraint on Schemas·](https://www.w3.org/TR/xmlschema-2/#dt-cos) and [·Validation Rule·](https://www.w3.org/TR/xmlschema-2/#dt-cvc) .

[must](https://www.w3.org/TR/xmlschema-2/#dt-must)

**must**

[non-numeric](https://www.w3.org/TR/xmlschema-2/#dt-non-numeric)

A datatype whose values are not [·numeric·](https://www.w3.org/TR/xmlschema-2/#dt-numeric) is said to be **non-numeric**.

[numeric](https://www.w3.org/TR/xmlschema-2/#dt-numeric)

A datatype is said to be **numeric** if its values are conceptually quantities (in some mathematical number system).

[order-relation](https://www.w3.org/TR/xmlschema-2/#dt-order-relation)

An **order relation** on a [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) is a mathematical relation that imposes a [·total order·](https://www.w3.org/TR/xmlschema-2/#dt-total-order) or a [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order) on the members of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[ordered](https://www.w3.org/TR/xmlschema-2/#dt-ordered)

A [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space), and hence a datatype, is said to be **ordered** if there exists an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) defined for that [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space).

[partial order](https://www.w3.org/TR/xmlschema-2/#dt-partial-order)

A **partial order** is an [·order-relation·](https://www.w3.org/TR/xmlschema-2/#dt-order-relation) that is **irreflexive**, **asymmetric** and **transitive**.

[primitive](https://www.w3.org/TR/xmlschema-2/#dt-primitive)

**Primitive** datatypes are those that are not defined in terms of other datatypes; they exist _ab initio_.

[regular expression](https://www.w3.org/TR/xmlschema-2/#dt-regex)

A **regular expression** is composed from zero or more [·branch·](https://www.w3.org/TR/xmlschema-2/#dt-branch)es, separated by `|` characters.

[restriction](https://www.w3.org/TR/xmlschema-2/#dt-restriction)

A datatype is said to be [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) by **restriction** from another datatype when values for zero or more [·constraining facet·](https://www.w3.org/TR/xmlschema-2/#dt-constraining-facet)s are specified that serve to constrain its [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space) and/or its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space) to a subset of those of its [·base type·](https://www.w3.org/TR/xmlschema-2/#dt-basetype).

[Schema Representation Constraint](https://www.w3.org/TR/xmlschema-2/#dt-src)

**Schema Representation Constraint**

[total order](https://www.w3.org/TR/xmlschema-2/#dt-total-order)

A **total order** is an [·partial order·](https://www.w3.org/TR/xmlschema-2/#dt-partial-order) such that for no _a_ and _b_ is it the case that _a <> b_.

[union](https://www.w3.org/TR/xmlschema-2/#dt-union)

**Union** datatypes are those whose [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s are the union of the [·value space·](https://www.w3.org/TR/xmlschema-2/#dt-value-space)s and [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space)s of one or more other datatypes.

[user-derived](https://www.w3.org/TR/xmlschema-2/#dt-user-derived)

**User-derived** datatypes are those [·derived·](https://www.w3.org/TR/xmlschema-2/#dt-derived) datatypes that are defined by individual schema designers.

[Validation Rule](https://www.w3.org/TR/xmlschema-2/#dt-cvc)

**Validation Rule**

[value space](https://www.w3.org/TR/xmlschema-2/#dt-value-space)

A **value space** is the set of values for a given datatype. Each value in the **value space** of a datatype is denoted by one or more literals in its [·lexical space·](https://www.w3.org/TR/xmlschema-2/#dt-lexical-space).

## H References

### [![next sub-section](https://www.w3.org/TR/xmlschema-2/next.jpg)](https://www.w3.org/TR/xmlschema-2/#non-normative-biblio)H.1 Normative

Clinger, WD (1990)

William D Clinger. _How to Read Floating Point Numbers Accurately._ In _Proceedings of Conference on Programming Language Design and Implementation_, pages 92-101. Available at: [ftp://ftp.ccs.neu.edu/pub/people/will/howtoread.ps](ftp://ftp.ccs.neu.edu/pub/people/will/howtoread.ps)

IEEE 754-1985

IEEE. _IEEE Standard for Binary Floating-Point Arithmetic._ See [http://standards.ieee.org/reading/ieee/std_public/description/busarch/754-1985_desc.html](http://standards.ieee.org/reading/ieee/std_public/description/busarch/754-1985_desc.html)

Namespaces in XML

World Wide Web Consortium. _Namespaces in XML_. Available at: [http://www.w3.org/TR/1999/REC-xml-names-19990114/](https://www.w3.org/TR/1999/REC-xml-names-19990114/)

RFC 1766

H. Alvestrand, ed. _RFC 1766: Tags for the Identification of Languages_ 1995. Available at: [http://www.ietf.org/rfc/rfc1766.txt](http://www.ietf.org/rfc/rfc1766.txt)

RFC 2045

N. Freed and N. Borenstein. _RFC 2045: Multipurpose Internet Mail Extensions (MIME) Part One: Format of Internet Message Bodies_. 1996. Available at: [http://www.ietf.org/rfc/rfc2045.txt](http://www.ietf.org/rfc/rfc2045.txt)

RFC 2396

Tim Berners-Lee, et. al. _RFC 2396: Uniform Resource Identifiers (URI): Generic Syntax._. 1998. Available at: [http://www.ietf.org/rfc/rfc2396.txt](http://www.ietf.org/rfc/rfc2396.txt)

RFC 2732

_RFC 2732: Format for Literal IPv6 Addresses in URL's_. 1999. Available at: [http://www.ietf.org/rfc/rfc2732.txt](http://www.ietf.org/rfc/rfc2732.txt)

RFC 3066

H. Alvestrand, ed. _RFC 3066: Tags for the Identification of Languages_ 1995. Available at: [http://www.ietf.org/rfc/rfc3066.txt](http://www.ietf.org/rfc/rfc3066.txt)

Unicode Database

The Unicode Consortium. _The Unicode Character Database_. Available at: [http://www.unicode.org/Public/3.1-Update/UnicodeCharacterDatabase-3.1.0.html](http://www.unicode.org/Public/3.1-Update/UnicodeCharacterDatabase-3.1.0.html)

XML 1.0 (Second Edition)

World Wide Web Consortium. _Extensible Markup Language (XML) 1.0, Second Edition._ Available at: [http://www.w3.org/TR/2000/WD-xml-2e-20000814](https://www.w3.org/TR/2000/WD-xml-2e-20000814)

XML Base

World Wide Web Consortium. XML Base. Available at: [http://www.w3.org/TR/2001/REC-xmlbase-20010627/](https://www.w3.org/TR/2001/REC-xmlbase-20010627/)

XML Linking Language

World Wide Web Consortium. XML Linking Language (XLink). Available at: [http://www.w3.org/TR/2001/REC-xlink-20010627/.](https://www.w3.org/TR/2001/REC-xlink-20010627/) Note: only the URI reference escaping procedure defined in Section 5.4 is normatively referenced.

XML Schema Part 1: Structures

XML Schema Part 1: Structures. Available at: [http://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html](https://www.w3.org/TR/2004/REC-xmlschema-1-20041028/structures.html)

XML Schema Requirements

World Wide Web Consortium. XML Schema Requirements. Available at: [http://www.w3.org/TR/1999/NOTE-xml-schema-req-19990215](https://www.w3.org/TR/1999/NOTE-xml-schema-req-19990215)

###  [![previous sub-section](https://www.w3.org/TR/xmlschema-2/previous.jpg)](https://www.w3.org/TR/xmlschema-2/#normative-biblio)H.2 Non-normative

Character Model

Martin J. Dürst and François Yergeau, eds. Character Model for the World Wide Web. World Wide Web Consortium Working Draft. 2001. Available at: [http://www.w3.org/TR/2001/WD-charmod-20010126/](https://www.w3.org/TR/2001/WD-charmod-20010126/)

Gay, DM (1990)

David M. Gay. _Correctly Rounded Binary-Decimal and Decimal-Binary Conversions._ AT&T Bell Laboratories Numerical Analysis Manuscript 90-10, November 1990. Available at: [http://cm.bell-labs.com/cm/cs/doc/90/4-10.ps.gz](http://cm.bell-labs.com/cm/cs/doc/90/4-10.ps.gz)

HTML 4.01

World Wide Web Consortium. Hypertext Markup Language, version 4.01. Available at: [http://www.w3.org/TR/1999/REC-html401-19991224/](https://www.w3.org/TR/1999/REC-html401-19991224/)

IETF INTERNET-DRAFT: IRIs

M. Dürst and M. Suignard . _Internationalized Resource Identifiers_ 2002. Available at: [http://www.w3.org/International/iri-edit/draft-duerst-iri-04.txt](https://www.w3.org/International/iri-edit/draft-duerst-iri-04.txt)

International Earth Rotation Service (IERS)

International Earth Rotation Service (IERS). See [http://maia.usno.navy.mil](http://maia.usno.navy.mil)

ISO 11404

ISO (International Organization for Standardization). _Language-independent Datatypes._ See [http://www.iso.ch/cate/d19346.html](http://www.iso.ch/cate/d19346.html)

ISO 8601

ISO (International Organization for Standardization). _Representations of dates and times, 1988-06-15._

ISO 8601:1998 Draft Revision

ISO (International Organization for Standardization). _Representations of dates and times, draft revision, 1998._

ISO 8601:2000 Second Edition

ISO (International Organization for Standardization). _Representations of dates and times, second edition, 2000-12-15._

Perl

The Perl Programming Language. See [http://www.perl.com/pub/language/info/software.html](http://www.perl.com/pub/language/info/software.html)

RDF Schema

World Wide Web Consortium. _RDF Schema Specification._ Available at: [http://www.w3.org/TR/2000/CR-rdf-schema-20000327/](https://www.w3.org/TR/2000/CR-rdf-schema-20000327/)

Ruby

World Wide Web Consortium. Ruby Annotation. Available at: [http://www.w3.org/TR/2001/WD-ruby-20010216/](https://www.w3.org/TR/2001/WD-ruby-20010216/)

SQL

ISO (International Organization for Standardization). _ISO/IEC 9075-2:1999, Information technology --- Database languages --- SQL --- Part 2: Foundation (SQL/Foundation)_. [Geneva]: International Organization for Standardization, 1999. See [http://www.iso.ch/cate/d26197.html](http://www.iso.ch/cate/d26197.html)

U.S. Naval Observatory Time Service Department

_Information about Leap Seconds_ Available at: [http://tycho.usno.navy.mil/leapsec.990505.html](http://tycho.usno.navy.mil/leapsec.990505.html)

Unicode Regular Expression Guidelines

Mark Davis. _Unicode Regular Expression Guidelines_, 1988. Available at: [http://www.unicode.org/unicode/reports/tr18/](http://www.unicode.org/unicode/reports/tr18/)

XML Schema Language: Part 0 Primer

World Wide Web Consortium. XML Schema Language: Part 0 Primer. Available at: [http://www.w3.org/TR/2004/REC-xmlschema-0-20041028/primer.html](https://www.w3.org/TR/2004/REC-xmlschema-0-20041028/primer.html)

XSL

World Wide Web Consortium. _Extensible Stylesheet Language (XSL)._ Available at: [http://www.w3.org/TR/2000/CR-xsl-20001121/](https://www.w3.org/TR/2000/CR-xsl-20001121/)

## I Acknowledgements (non-normative)

The following have contributed material to the first edition of this specification:

> Asir S. Vedamuthu, webMethods, Inc
> Mark Davis, IBM

Co-editor Ashok Malhotra's work on this specification from March 1999 until February 2001 was supported by IBM. From February 2001 until May 2004 it was supported by Microsoft.

The editors acknowledge the members of the XML Schema Working Group, the members of other W3C Working Groups, and industry experts in other forums who have contributed directly or indirectly to the process or content of creating this document. The Working Group is particularly grateful to Lotus Development Corp. and IBM for providing teleconferencing facilities.

At the time the first edition of this specification was published, the members of the XML Schema Working Group were:

- Jim Barnette, Defense Information Systems Agency (DISA)
- Paul V. Biron, Health Level Seven
- Don Box, DevelopMentor
- Allen Brown, Microsoft
- Lee Buck, TIBCO Extensibility
- Charles E. Campbell, Informix
- Wayne Carr, Intel
- Peter Chen, Bootstrap Alliance and LSU
- David Cleary, Progress Software
- Dan Connolly, W3C (_staff contact_)
- Ugo Corda, Xerox
- Roger L. Costello, MITRE
- Haavard Danielson, Progress Software
- Josef Dietl, Mozquito Technologies
- David Ezell, Hewlett-Packard Company
- Alexander Falk, Altova GmbH
- David Fallside, IBM
- Dan Fox, Defense Logistics Information Service (DLIS)
- Matthew Fuchs, Commerce One
- Andrew Goodchild, Distributed Systems Technology Centre (DSTC Pty Ltd)
- Paul Grosso, Arbortext, Inc
- Martin Gudgin, DevelopMentor
- Dave Hollander, Contivo, Inc (_co-chair_)
- Mary Holstege, Invited Expert
- Jane Hunter, Distributed Systems Technology Centre (DSTC Pty Ltd)
- Rick Jelliffe, Academia Sinica
- Simon Johnston, Rational Software
- Bob Lojek, Mozquito Technologies
- Ashok Malhotra, Microsoft
- Lisa Martin, IBM
- Noah Mendelsohn, Lotus Development Corporation
- Adrian Michel, Commerce One
- Alex Milowski, Invited Expert
- Don Mullen, TIBCO Extensibility
- Dave Peterson, Graphic Communications Association
- Jonathan Robie, Software AG
- Eric Sedlar, Oracle Corp.
- C. M. Sperberg-McQueen, W3C (_co-chair_)
- Bob Streich, Calico Commerce
- William K. Stumbo, Xerox
- Henry S. Thompson, University of Edinburgh
- Mark Tucker, Health Level Seven
- Asir S. Vedamuthu, webMethods, Inc
- Priscilla Walmsley, XMLSolutions
- Norm Walsh, Sun Microsystems
- Aki Yoshida, SAP AG
- Kongyi Zhou, Oracle Corp.

The XML Schema Working Group has benefited in its work from the participation and contributions of a number of people not currently members of the Working Group, including in particular those named below. Affiliations given are those current at the time of their work with the WG.

- Paula Angerstein, Vignette Corporation
- David Beech, Oracle Corp.
- Gabe Beged-Dov, Rogue Wave Software
- Greg Bumgardner, Rogue Wave Software
- Dean Burson, Lotus Development Corporation
- Mike Cokus, MITRE
- Andrew Eisenberg, Progress Software
- Rob Ellman, Calico Commerce
- George Feinberg, Object Design
- Charles Frankston, Microsoft
- Ernesto Guerrieri, Inso
- Michael Hyman, Microsoft
- Renato Iannella, Distributed Systems Technology Centre (DSTC Pty Ltd)
- Dianne Kennedy, Graphic Communications Association
- Janet Koenig, Sun Microsystems
- Setrag Khoshafian, Technology Deployment International (TDI)
- Ara Kullukian, Technology Deployment International (TDI)
- Andrew Layman, Microsoft
- Dmitry Lenkov, Hewlett-Packard Company
- John McCarthy, Lawrence Berkeley National Laboratory
- Murata Makoto, Xerox
- Eve Maler, Sun Microsystems
- Murray Maloney, Muzmo Communication, acting for Commerce One
- Chris Olds, Wall Data
- Frank Olken, Lawrence Berkeley National Laboratory
- Shriram Revankar, Xerox
- Mark Reinhold, Sun Microsystems
- John C. Schneider, MITRE
- Lew Shannon, NCR
- William Shea, Merrill Lynch
- Ralph Swick, W3C
- Tony Stewart, Rivcom
- Matt Timmermans, Microstar
- Jim Trezzo, Oracle Corp.
- Steph Tryphonas, Microstar

The lists given above pertain to the first edition. At the time work on this second edition was completed, the membership of the Working Group was:

- Leonid Arbouzov, Sun Microsystems
- Jim Barnette, Defense Information Systems Agency (DISA)
- Paul V. Biron, Health Level Seven
- Allen Brown, Microsoft
- Charles E. Campbell, Invited expert
- Peter Chen, Invited expert
- Tony Cincotta, NIST
- David Ezell, National Association of Convenience Stores
- Matthew Fuchs, Invited expert
- Sandy Gao, IBM
- Andrew Goodchild, Distributed Systems Technology Centre (DSTC Pty Ltd)
- Xan Gregg, Invited expert
- Mary Holstege, Mark Logic
- Mario Jeckle, DaimlerChrysler
- Marcel Jemio, Data Interchange Standards Association
- Kohsuke Kawaguchi, Sun Microsystems
- Ashok Malhotra, Invited expert
- Lisa Martin, IBM
- Jim Melton, Oracle Corp
- Noah Mendelsohn, IBM
- Dave Peterson, Invited expert
- Anli Shundi, TIBCO Extensibility
- C. M. Sperberg-McQueen, W3C (_co-chair_)
- Hoylen Sue, Distributed Systems Technology Centre (DSTC Pty Ltd)
- Henry S. Thompson, University of Edinburgh
- Asir S. Vedamuthu, webMethods, Inc
- Priscilla Walmsley, Invited expert
- Kongyi Zhou, Oracle Corp.

We note with sadness the accidental death of Mario Jeckle shortly after the completion of work on this document. In addition to those named above, several people served on the Working Group during the development of this second edition:

- Oriol Carbo, University of Edinburgh
- Tyng-Ruey Chuang, Academia Sinica
- Joey Coyle, Health Level 7
- Tim Ewald, DevelopMentor
- Nelson Hung, Corel
- Melanie Kudela, Uniform Code Council
- Matthew MacKenzie, XML Global
- Cliff Schmidt, Microsoft
- John Stanton, Defense Information Systems Agency
- John Tebbutt, NIST
- Ross Thompson, Contivo
- Scott Vorthmann, TIBCO Extensibility
