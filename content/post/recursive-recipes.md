
---
title: "Recursive recipes"
date: 2018-05-29T05:40:05-07:00
tags: [food]
slug: recursive-recipes
description: The story of a website that can determine recipes from scratch, based on the recipes for the ingredients.
image: kjkl
written: ["2018-05-29","2018-05","2018"]
---

I've thought a lot about the recipe when making food.[^food] Lately I've been thinking about what kind of recipe you need to truly make something *from scratch*. To truly make a recipe *from scratch* you need only a few ingredients: sun, water, soil, and seeds.[^seeds] I previously made a website with various recipes that lets you see the work flows for creating different foods from scratch.[^1] I've now created another website, **[recursive.recipes](https://recursive.recipes)**,  which lets you scale recipes and better visualize their *costs* (time + money) with creating them from scratch various stages.


[![Visualization of the network for making chocolate chip cookies from scratch.](/img/chocolate_chip_cookies_graph.png)](https://recursive.recipes/recipe/chocolate-chip-cookies?amount=36&timelimit=37100&ingredientsToBuild=)


Making this tool was a little challenging, but in a fun way. The costs in time turned out to be non-trivial. I determined that there needed to be two types of time costs - *parallel* and *serial*. Basically, the *serial* time scales with the quantity, while *parallel* time does not (i.e. it can be done in parallel). For example, when harvesting a field of wheat, the parallel time is the time it takes to grow, which is independent of the quantity. The serial time is the time it takes to harvest, which is dependent on the quantity.

The monetary costs also were not easy to determine. I used the latest grocery prices, and tried to scale everything to volumetric or whole measures, for comparing. Obviously this becomes problematic when trying to determine how much volume a 1 lb steak takes up as it also requires computing the average density of the substance. Also the costs themselves were determined for farms from farming data, although they can have quite a lot of variance.

The hardest part where more questions about one-time cost versus recurring cost. The calculator I was making calculated one-time costs. But then, the advantages of farming are based on recurring costs. That is, you don't buy an entire cow just to make milk for one stick of butter. You obviously would use the cow over and over again to create many sticks of butter. How many? This is the hard thing to find out. Also, this might depend on whether you are making butter or just pasteurizing to make milk. For now, though, I've assumed the things like this that have huge recurring benefits to have a cost of free.

There is still room for improvement. All of the data for the tool is in a single file, [recipes.toml](https://github.com/schollz/recursive-recipes/blob/master/recipes.toml), so it is readily amended. Actually, the entire website is open-source[^website] and on [Github](https://github.com/schollz/recursive-recipes) if you are interested in checking it out or making a pull request for your favorite recipe.

[^food]: Definitely too much. I've already made a website for [pair-wise ingredient comparison](http://www.foodpairing.ninja/) and for [averaging recipes](https://www.averagecookbook.com/).
[^seeds]: By seeds I mean from animals (eggs), bacteria (cells), fungi (spores), or plants.
[^1]: [timetomakefood.com](https://timetomakefood.com/loaf-of-bread/) - for more information on my previous tool, check out a post I did previously titled *"[Recursive ingredients](/recursive-ingredients)"*.
[^website]: Technically, all websites are open-source, but here I mean you can access the backend as well.