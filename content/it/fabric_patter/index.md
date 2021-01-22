---
title: "Паттерн Фабрика и Абстрактная Фабрика"
date: 2020-04-07T21:03:21+03:00
type: it
draft: false
linkhabr: false
description: Паттерн Фабрика и Абстрактная Фабрика Python ООП
image: default_open_graph.jpg
tags: [python]
---

Пара строчек о паттернах Фабрика и Абстрактная фабрика на Python.

Как известно бандой четырех Фабрика и Абстрактная фабрика были отнесены к порождающим паттернам проектирования.
Цель этих методов решить проблему создания однотипных объектов путем создания общих метаклассов, позволяя скрыть
сложную механику через общие интерфейсы.

<!--more-->
Для начала давайте попробуем написать Фабрику. 
Соль этого метода в том, чтобы сделать метакласс, который будет создавать объекты в зависимости от переданных параметров,
тем самым инкапсулирует часть методов для создания объектов и унифицировать доступ к ним.

Давайте напишим наш класс:
{{< highlight python  >}}
    class FabrickСreature:
        """ Фабричный Класс который создает животное иcходя из переданного name"""
        _subclasses = []
        
        def __init_subclass__(cls):
            FabrickСreature._subclasses.append(cls)
        
        def getСreature(self):
            return self.creature
        
        def create(self, name, type, count=1):
            for i in self._subclasses:
                if name == i.__name__:
                    self.name = name
                    self.creature = i(count, type)
                    return
                print(f"нет таких животных как {name}")
        
        def evening(self):
            """что было вечером"""
            print(f"Что было вечером с {self.name} - {self.creature}")
            self.creature.in_the_evening()
{{< /highlight >}}
<br/>

Обычно Сначала создают фабрику с общими абстрактными методами а потом от нее наследую Фабричные классы для каждого
типа объектов.

В нашем случае же мы чуть чуть схитрим и воспользовавшись возможностями Python в виде self._subclasses и
множественного наследования <em> ( вы увидите ниже как классы все наследуют класс FabrickСreature) </em>. Тем самым
мы избежим множественных if-ов и чуточку упростим код.

Также хочу отметить метод evening.
Класс фабрика не только создает объекты но также очень часто имеет бизнес логику для запуска созданных им объектов.

Теперь напишим наших Животных.
{{< highlight python  >}}
    class Creature():
        def __init__(self, c, t):
            self.count = c
            self.type = t
        
        def in_the_evening(self):
            """Действия животного вечером"""
            self.go()
            self.eat()
            self.sleep()
        
        def __repr__(self):
            return f"< {self.type} >"
        
        def eat(self):
            """Абстрактный метод для еды"""
        
        def go(self):
            """Абстрактный метод двигаться"""
        
        def sleep(self):
            """Абстрактный метод для сна"""
    
    
    class Fish(FabrickСreature, Creature):
        """Рыбы"""
        
        def getName(self):
            return self.name
        
        def eat(self):
            print("Рыба ела водоросли")
        
        def go(self):
            print("Косяк рыб подплыл к кусту где сидели птички")
        
        def sleep(self):
            print("Рыбы не спят")
    
    
    class Bird(FabrickСreature, Creature):
        """Птица"""
        
        def eat(self):
            print(f"Птички в количестве {self.count} клевали почки на дереве")
        
        def go(self):
            print(f"Птички в количестве {self.count} прилетели к кусту возле берега")
        
        def sleep(self):
            print("птички наелись и заснули")
    
    
    class Cat(FabrickСreature, Creature):
        """Кошка"""
        
        def eat(self):
            print("Кошка съела пойманную птичку")
        
        def go(self):
            print("Кошка подобралась к спящей птичке и напала")
        
        def sleep(self):
            print("Кошка пошла спать домой")

{{< /highlight >}}
<br/>
Как вы видите каждый объект Животного мы унаследовали от двух объектов Creature в котором описаны абстрактные
методы, Также мы унаследовали класс FabrickСreature, тем самым мы зарегистрировали все объекты животных на фабрике.

Теперь можно вызывать:

{{< highlight python  >}}
fabric = FabrickСreature()
fabric.create("Bird", "Воробей", 5)
fabric.evening()
print(f"  {fabric.getСreature()}", "\n")

fabric.create("Fish", "Карасики")
fabric.evening()
print(f" {fabric.getСreature()}", "\n")

fabric.create("Cat", "Муська")
fabric.evening()
print(f"  {fabric.getСreature()}", "\n")
{{< /highlight >}}
<br/>

C Фабричным методом допустим понятно. Но в чем же суть Абстрактной фабрики.

Основную мысль этого метода можно описать так.

Допустим из примера выше у нас есть животные не простые а вымышленные.
И в связи с этим Фабрика для обычных животных нам может не подойти,
из за того что способы создания этого животного может изменится.

Так вот абстрактная фабрика это как бы мама всех фабрик в которой описываются основные
методы работы фабрик для того чтобы можно было ими пользоваться. Но как эти методы
работают как они создают этих животных. Это уже должно упасть на откуп конкретному типу фабрики.

Для начала создадим Абстрактную фабрику и конкретные фабрики с разными способами создания животных.

{{< highlight python  >}}
class AbstractFabrickСreature():
    
    def create(self):
        """Абстрактный метод """
    
    def evening(self):
        """Абстрактный метод """
    
    def getСreature(self):
        """Абстрактный метод """


class FabrickСreature(AbstractFabrickСreature):
    """ Фабричный Класс который создает животное иcходя из переданного name"""
    _subclasses = []

    def __init_subclass__(cls):
        FabrickСreature._subclasses.append(cls)
    
    def getСreature(self):
        return self.creature

    def create(self, name, type, count=1):
        for i in self._subclasses:
            if name == i.__name__:
            self.name = name
            self.creature = i(count, type)
            return
        print(f"нет таких животных как {name}")

    def evening(self):
        """что было как то вечером"""
        print(f"Что было вечером с {self.name} - {self.creature}")
        self.creature.in_the_evening()


    class FabrickMagicСreature(AbstractFabrickСreature):
        """ Фабричный Класс который создает ВОЛШЕБНЫХ животное иcходя из переданного name"""
    
    def getСreature(self):
        return self.creature

    def create(self, name, type):
        if name == "Unicorn":
            self.creature = Unicorn(type)
        elif name == "Hedgehog":
            self.creature = Hedgehog(type)
        else:
            print("Нет таких волшебных животных")
            exit()

    def evening(self):
        """что было как то вечером"""
        print(f"Что было вечером  В ВОЛШЕБНОМ ЛЕСУ ! ! ! ! !  ! ! ")
        self.creature.in_the_evening()
{{< /highlight >}}
<br/>


Как видите у нас есть 2 конкретные фабрики которые по разному реализуют функции создания объектов
но имеют общий интерфейс через базовый класс AbstractFabrickСreature.

Теперь создадим Волшебных животных:


{{< highlight python  >}}
    class MagicCreature():
        def __init__(self, t):
        self.type = t
    
    def in_the_evening(self):
        self.go()
        self.see()
    
    def __repr__(self):
        return f"< {self.type} >"
    
    def see(self):
        """Абстрактный метод для Взгляда"""
        
    def go(self):
        """Абстрактный метод двигаться"""
    
    
    class Hedgehog(MagicCreature):
        def see(self):
        print("Ежик смотрел в Туман и искал Единорожку")
    
    def go(self):
        print("Ежик взяв узелок с яблоками пошел на опушку леса")   
    
    
    class Unicorn(MagicCreature):
        def see(self):
        print("Единорог виде только макушку ежика")
        
    def go(self):
        print("Единорого стоял на опушке леса")
{{< /highlight >}}
<br/>
В целом по Абстрактному паттерну все. ))

С листингами кода можно ознакомится по ссылкам:

<ul class="">
    <li><a href="https://gist.github.com/IzyI/a0952fa8e63cc1d22e511bee2cdf923f" target="_blank">Паттерн Фабрика</a></li>
    <li><a href="https://gist.github.com/IzyI/5726502eca6c988aa11aa37ad4cda4cc" target="_blank">Паттерн Абстрактная
        Фабрика</a></li>
</ul>

