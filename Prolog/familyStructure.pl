/* :- initialization(main). */
/* main :- write('Hello World!'), nl, halt. */
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
is_mother(jessicam, adam).
is_father(michael, adam).

is_brother(michael, mary).
is_sister(mary, michael).
