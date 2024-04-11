#include common_scripts\utility;
#include maps\mp\gametypes_zm\_hud_util;
#include maps\mp\zombies\_zm_utility;

test_func()
{
	name = "Max";
	age = 22;
	age = 30;
	str = name + ", " + age;
	return name + ": " + age;
}

hello(name)
{
	print("Hello, "+name);
	print(arg1,arg2);
	print();
}

func_with_multiple_args(arg1, arg2, arg3)
{
	result = arg1 + arg2 + arg3;
	return result;
}
