male(john).
male(michael).
male(adam).
female(anne).
female(jessica).
female(mary).

is_spouse(anne, john).
is_spouse(michael, jessica).

is_mother(anne, michael).
is_father(john, michael).
is_mother(anne, mary).
is_father(john, mary).
is_mother(jessica, adam).
is_father(michael, adam).

is_brother(michael, mary).
is_sister(mary, michael).

is_grandmother(anne, adam).
is_grandfather(john, adam).
is_grandmother(anne, sam).
is_grandfather(john, sam).

is_aunt(jessica, sam).
is_uncle(michael, sam).
is_aunt(mary, adam).
is_uncle(phil, adam).

is_cousin(adam, sam).
is_cousin(sam, adam).
