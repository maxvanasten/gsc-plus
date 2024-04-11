#include common_scripts\utility;
#include maps\mp\gametypes_zm\_hud_util;
#include maps\mp\zombies\_zm_utility;

get_hello_string(name)
{
	first_part = "Hello, ";
	second_part = name + "!";
	return first_part + second_part;
}

hello_on_obj()
{
	print("Hello, "+self.name+"!");
}

greet_person(person)
{
	person hello_on_obj();
	person thread hello_on_obj();
}
