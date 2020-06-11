package ustatynamegenerator

import (
	"fmt"
	"math/rand"
)

var (
	left = [...]string{
		"admiring",
		"adoring",
		"affectionate",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"beautiful",
		"blissful",
		"bold",
		"boring",
		"brave",
		"busy",
		"charming",
		"clever",
		"cool",
		"compassionate",
		"competent",
		"condescending",
		"confident",
		"cranky",
		"crazy",
		"dazzling",
		"dwarf",
		"determined",
		"distracted",
		"dreamy",
		"eager",
		"ecstatic",
		"elastic",
		"elated",
		"elegant",
		"eloquent",
		"epic",
		"exciting",
		"fervent",
		"festive",
		"flamboyant",
		"focused",
		"friendly",
		"frosty",
		"funny",
		"gallant",
		"gifted",
		"goofy",
		"gracious",
		"great",
		"happy",
		"hardcore",
		"heuristic",
		"hopeful",
		"hungry",
		"infallible",
		"inspiring",
		"interesting",
		"intelligent",
		"jolly",
		"jovial",
		"keen",
		"kind",
		"laughing",
		"loving",
		"lucid",
		"magical",
		"mystifying",
		"modest",
		"musing",
		"naughty",
		"nervous",
		"nice",
		"nifty",
		"nostalgic",
		"objective",
		"optimistic",
		"peaceful",
		"pedantic",
		"pensive",
		"practical",
		"priceless",
		"quirky",
		"quizzical",
		"recursing",
		"relaxed",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"silly",
		"sleepy",
		"stoic",
		"strange",
		"stupefied",
		"suspicious",
		"sweet",
		"tender",
		"thirsty",
		"trusting",
		"unruffled",
		"upbeat",
		"vibrant",
		"vigilant",
		"vigorous",
		"wizardly",
		"wonderful",
		"xenodochial",
		"youthful",
		"zealous",
		"zen",
		"planet",
	}

	right = [...]string{
		//Ceres /ˈsɪəriːz/[16] (minor-planet designation: 1 Ceres) is the largest object in the main asteroid belt that lies between the orbits of Mars and Jupiter. With a diameter of 940 km (580 mi), Ceres is both the largest of the asteroids and the only unambiguous dwarf planet currently inside Neptune's orbit.[c] It is the 25th-largest body in the Solar System within the orbit of Neptune.[17] https://en.wikipedia.org/wiki/1_Ceres
		"ceres",

		//Pallas (minor-planet designation: 2 Pallas) is the second asteroid to have been discovered, after 1 Ceres. It is one of the largest asteroids in the Solar System, and is a likely remnant protoplanet. With an estimated 7% of the mass of the asteroid belt, it is the third-most-massive (and third-largest) asteroid, being three quarters the mass of 4 Vesta and one quarter the mass of Ceres. It is about 510 kilometers (320 mi) in diameter, slightly smaller than Vesta. https://en.wikipedia.org/wiki/2_Pallas
		"pallas",

		//3 Juno is a large asteroid in the asteroid belt. Juno was the third asteroid discovered, in 1804, by German astronomer Karl Harding.[15] It is the 11th-largest asteroid, and one of the two largest stony (S-type) asteroids, along with 15 Eunomia. It is estimated to contain 1% of the total mass of the asteroid belt.[16] https://en.wikipedia.org/wiki/3_Juno
		"juno",

		//Vesta (minor-planet designation: 4 Vesta) is one of the largest objects in the asteroid belt, with a mean diameter of 525 kilometres (326 mi).[9] It was discovered by the German astronomer Heinrich Wilhelm Matthias Olbers on 29 March 1807[7] and is named after Vesta, the virgin goddess of home and hearth from Roman mythology. https://en.wikipedia.org/wiki/4_Vesta
		"vesta",

		//Astraea /æˈstriːə/ (minor planet designation: 5 Astraea) is a large asteroid from the asteroid belt. Its surface is highly reflective (bright) and its composition is probably a mixture of nickel–iron with silicates of magnesium and iron. It is an S-type asteroid in the Tholen classification system.[3] https://en.wikipedia.org/wiki/5_Astraea
		"astrea",

		//Hebe /ˈhiːbiː/ (minor planet designation: 6 Hebe) is a large main-belt asteroid, containing around 0.5% of the mass of the belt. However, due to its apparently high bulk density (greater than that of the Moon or even Mars), Hebe does not rank among the top twenty asteroids by volume. This high bulk density suggests an extremely solid body that has not been impacted by collisions, which is not typical of asteroids of its size – they tend to be loosely-bound rubble piles. https://en.wikipedia.org/wiki/6_Hebe
		"hebe",

		//Iris (minor planet designation: 7 Iris) is a large main-belt asteroid orbiting the Sun between Mars and Jupiter. It is the fourth-brightest object in the asteroid belt. It is classified as an S-type asteroid, meaning that it has a stony composition. https://en.wikipedia.org/wiki/7_Iris
		"iris",

		//Flora (minor planet designation: 8 Flora) is a large, bright main-belt asteroid. It is the innermost large asteroid: no asteroid closer to the Sun has a diameter above 25 kilometres or two-elevenths that of Flora itself, and not until the tiny 149 Medusa was discovered was a single asteroid orbiting at a closer mean distance known.[7] It is the seventh-brightest asteroid with a mean opposition magnitude of +8.7.[8] Flora can reach a magnitude of +7.9 at a favorable opposition near perihelion, such as occurred in November 2007. Flora may be the residual core of an intensely heated, thermally evolved, and magmatically differentiated planetesimal which was subsequently disrupted.[9] https://en.wikipedia.org/wiki/8_Flora
		"flora",

		//Metis (minor planet designation: 9 Metis) is one of the larger main-belt asteroids. It is composed of silicates and metallic nickel-iron, and may be the core remnant of a large asteroid that was destroyed by an ancient collision.[8] Metis is estimated to contain just under half a percent of the total mass of the asteroid belt.[9] https://en.wikipedia.org/wiki/9_Metis
		"metis",

		//Hygiea (minor planet designation: 10 Hygiea) is a major asteroid located in the main asteroid belt. With a diameter of 434 kilometres (270 mi) and a mass estimated to be 2.9% of the total mass of the belt,[10] it is the fourth-largest asteroid in the Solar System by both volume and mass. In some spectral classifications it is the largest of the dark C-type asteroids with a carbonaceous surface, in others it is second after 1 Ceres. https://en.wikipedia.org/wiki/10_Hygiea
		"hygiea",

		//Parthenope[6] (minor planet designation: 11 Parthenope) is a large, bright main-belt asteroid. https://en.wikipedia.org/wiki/11_Parthenope
		"parthenope",

		//Victoria (minor planet designation: 12 Victoria) is a large main-belt asteroid, orbiting the Sun with a period of 3.56 years and an eccentricity of 0.221. It is a stony (S-type) asteroid, about 112–124 km across with an albedo of 0.18. Victoria has been observed to occult a star three times since its discovery. Radar and speckle interferometry observations show that the shape of Victoria is elongated, and it is suspected to be a binary asteroid, with a moon of irregular shape.[5] https://en.wikipedia.org/wiki/12_Victoria
		"victoria",

		//Egeria (minor planet designation: 13 Egeria) is a large main-belt G-type asteroid.[8] It was discovered by Annibale de Gasparis on November 2, 1850. Egeria was named by Urbain Le Verrier, whose computations led to the discovery of Neptune, after the mythological nymph Egeria of Aricia, Italy, the wife of Numa Pompilius, second king of Rome.[9] https://en.wikipedia.org/wiki/13_Egeria
		"egeria",

		//Irene /aɪˈriːniː/ (minor planet designation: 14 Irene) is a large main-belt asteroid, discovered by John Russell Hind on May 19, 1851. https://en.wikipedia.org/wiki/14_Irene
		"irene",

		//Eunomia (minor planet designation 15 Eunomia) is a very large asteroid in the inner asteroid belt. It is the largest of the stony (S-type) asteroids, and somewhere between the 8th- and 12th-largest main-belt asteroid overall (uncertainty in diameters causes uncertainty in its ranking). It is the largest Eunomian asteroid, and is estimated to contain 1% of the mass of the asteroid belt.[10][11] https://en.wikipedia.org/wiki/15_Eunomia
		"eunomia",

		//16 Psyche /ˈsaɪkiː/ is one of the most massive asteroids in the asteroid belt. This object is over 200 km (120 mi) in diameter and contains about 1% of the mass of the entire asteroid belt. It is thought to be the exposed iron core of a protoplanet,[8] and is the most massive metallic M-type asteroid. Psyche was discovered by the Italian astronomer Annibale de Gasparis on 17 March 1852 from Naples and named after the Greek mythological figure Psyche.[9] The prefix "16" signifies that it was the sixteenth minor planet in order of discovery. https://en.wikipedia.org/wiki/16_Psyche
		"psyche",

		//Thetis, minor planet designation 17 Thetis, is a stony asteroid from the inner regions of the asteroid belt, approximately 90 kilometers in diameter. It was discovered on 17 April 1852, by German astronomer Robert Luther at Bilk Observatory in Düsseldorf, Germany.[3] He named his first asteroid discovery after Thetis from Greek mythology.[2] https://en.wikipedia.org/wiki/17_Thetis
		"thetis",

		//Melpomene (minor planet designation: 18 Melpomene) is a large, bright main-belt asteroid that was discovered by J. R. Hind on June 24, 1852,[10] and named after Melpomenē, the Muse of tragedy in Greek mythology. It is classified as an S-type asteroid and is composed of silicates and metals. https://en.wikipedia.org/wiki/18_Melpomene
		"melpomene",

		//Fortuna (minor planet designation: 19 Fortuna) is one of the largest main-belt asteroids. It has a composition similar to 1 Ceres: a darkly colored surface that is heavily space-weathered with the composition of primitive organic compounds, including tholins. https://en.wikipedia.org/wiki/19_Fortuna
		"fortuna",

		//Massalia, minor planet designation 20 Massalia, is a stony asteroid and the parent body of the Massalia family located in the inner region of the asteroid belt, approximately 145 kilometers (90 miles) in diameter. Discovered by Italian astronomer Annibale de Gasparis on 19 September 1852, it was named for the French city of Marseille, from which the independent discover Jean Chacornac sighted it the following night.[2] https://en.wikipedia.org/wiki/20_Massalia
		"massalia",

		//Lutetia (minor planet designation: 21 Lutetia) is a large asteroid in the asteroid belt of an unusual spectral type. It measures about 100 kilometers in diameter (120 km along its major axis). It was discovered in 1852 by Hermann Goldschmidt, and is named after Lutetia, the Latin name of Paris. https://en.wikipedia.org/wiki/21_Lutetia
		"lutetia",

		//Kalliope[6] (minor planet designation: 22 Kalliope) is a large M-type asteroid from the asteroid belt discovered by J. R. Hind on 16 November 1852. It is named after Calliope, the Greek Muse of epic poetry. It is orbited by a small moon named Linus. https://en.wikipedia.org/wiki/22_Kalliope
		"kalliope",

		//Thalia[5] (minor planet designation: 23 Thalia) is a large main-belt asteroid. It was discovered by J. R. Hind on December 15, 1852, at the private observatory of W. Bishop, located in Hyde Park, London, England.[6] Bishop named it after Thalia, the Muse of comedy and pastoral poetry in Greek mythology.[7] https://en.wikipedia.org/wiki/23_Thalia
		"thalia",

		//Themis (minor planet designation: 24 Themis) is one of the largest asteroids in the asteroid belt. It is also the largest member of the Themistian family. It was discovered by Annibale de Gasparis on 5 April 1853. It is named after Themis, the personification of natural law and divine order in Greek mythology. https://en.wikipedia.org/wiki/24_Themis
		"themis",

		//Phocaea /foʊˈsiːə/ (minor planet designation: 25 Phocaea) is a stony asteroid from the inner regions of the asteroid belt, approximately 75 kilometers in diameter. It is the parent body of the Phocaea family. Discovered by Jean Chacornac in 1853, it was named after the ancient Greek city of Phocaea. https://en.wikipedia.org/wiki/25_Phocaea
		"phocaea",

		//Proserpina (minor planet designation: 26 Proserpina) is a main-belt asteroid discovered by R. Luther on May 5, 1853. It is named after the Roman goddess Proserpina, the daughter of Ceres and the Queen of the Underworld. Another main-belt asteroid, 399 Persephone, discovered in 1895, is named after her Greek counterpart. https://en.wikipedia.org/wiki/26_Proserpina
		"proserpina",

		//Euterpe, minor planet designation 27 Euterpe, is a stony asteroid and parent body of the Euterpe family, located in the inner asteroid belt, approximately 100 kilometers in diameter. It was discovered by English astronomer John Russell Hind at George Bishop's Observatory in London on 8 November 1853. The asteroid was named after Euterpe, the Muse of music in Greek mythology.[2][22] https://en.wikipedia.org/wiki/27_Euterpe
		"euterpe",

		//Bellona (minor planet designation: 28 Bellona) is a large main-belt asteroid. It was discovered by R. Luther on March 1, 1854, and named after Bellōna, the Roman goddess of war; the name was chosen to mark the beginning of the Crimean War. https://en.wikipedia.org/wiki/28_Bellona
		"bellona",

		//Amphitrite (minor planet designation: 29 Amphridite) is one of the largest S-type asteroids, approximately 200 kilometers (124 miles) in diameter, and probably third largest after Eunomia and Juno, although Iris and Herculina are similar in size. https://en.wikipedia.org/wiki/29_Amphitrite
		"amphitrite",

		//Urania (minor planet designation: 30 Urania) is a large main-belt asteroid that was discovered by English astronomer John Russell Hind on July 22, 1854.[1] It was his last asteroid discovery. This object is named after Urania, the Greek Muse of astronomy. Initial orbital elements for 30 Urania were published by Wilhelm Günther, an assistant at Breslau Observatory.[6] https://en.wikipedia.org/wiki/30_Urania
		"urania",

		//Euphrosyne (minor planet designation: 31 Euphrosyne) is perhaps the 12th-largest asteroid in the asteroid belt, and possibly one of the half-dozen or so most massive. It was discovered by James Ferguson on September 1, 1854, the first asteroid found from North America. It is named after Euphrosyne, one of the Charites in Greek mythology. In 2019 a small companion was discovered. https://en.wikipedia.org/wiki/31_Euphrosyne
		"euphrosyne",

		//Pomona (minor planet designation: 32 Pomona) is a stony main-belt asteroid that is 81 km in diameter. It was discovered by German-French astronomer Hermann Mayer Salomon Goldschmidt on October 26, 1854,[3] and is named after Pōmōna, the Roman goddess of fruit trees. https://en.wikipedia.org/wiki/32_Pomona
		"pomona",

		//Polyhymnia (minor planet designation: 33 Polyhymnia) is a main belt asteroid that was discovered by French astronomer Jean Chacornac on October 28, 1854[1] and named after Polyhymnia, the Greek Muse of sacred hymns. https://en.wikipedia.org/wiki/33_Polyhymnia
		"polyhymnia",

		//Circe, minor planet designation 34 Circe, is a large, very dark main-belt asteroid. It was discovered by French astronomer J. Chacornac on April 6, 1855, and named after Circe, the bewitching queen of Aeaea island in Greek mythology. https://en.wikipedia.org/wiki/34_Circe
		"circe",

		//Leukothea (minor planet designation: 35 Leukothea) is a large, dark asteroid from the asteroid belt It was discovered by German astronomer Karl Theodor Robert Luther on April 19, 1855,[5] and named after Leukothea, a sea goddess in Greek mythology. 35 Leukothea is a C-type asteroid in the Tholen classification system.[2] https://en.wikipedia.org/wiki/35_Leukothea
		"leukothea",

		//Atalante (minor planet designation: 36 Atalante) is a large, dark main-belt asteroid. It was discovered by the German-French astronomer H. Goldschmidt on October 5, 1855, and named by French mathematician Urbain Le Verrier after the Greek mythological heroine Atalanta (of which Atalante is the French and German form, pronounced nearly the same as 'Atalanta' in English).[6] It was rendered 'Atalanta' in English sources in the 19th century.[2] The asteroid is also classified as a C-type one, according to the Tholen classification system.[1] https://en.wikipedia.org/wiki/36_Atalante
		"atalante",

		//Fides /ˈfaɪdiːz/ (minor planet designation: 37 Fides) is a large main-belt asteroid. It was discovered by German astronomer Karl Theodor Robert Luther on October 5, 1855,[5] and named after Fides, the Roman goddess of loyalty. Fides was the last of the main-belt asteroids to be assigned an iconic symbol.[6] 37 Fides is also a S-type asteroid in the Tholen classification system.[2] https://en.wikipedia.org/wiki/37_Fides
		"fides",

		//Leda (minor planet designation: 38 Leda) is a large, dark main-belt asteroid that was discovered by French astronomer J. Chacornac on January 12, 1856, and named after Leda, the mother of Helen of Troy in Greek mythology. In the Tholen classification system, it is categorized as a carbonaceous C-type asteroid, while the Bus asteroid taxonomy system lists it as a Cgh asteroid.[4] The spectra of the asteroid displays evidence of aqueous alteration.[5] https://en.wikipedia.org/wiki/38_Leda
		"leda",

		//Laetitia (minor planet designation: 39 Laetitia) is a large main-belt asteroid that was discovered by French astronomer Jean Chacornac on 9 February 1856[8] and named after "Laetitia", one of the epithets of Ceres, Roman goddess of fertility and abundance. Photometric observations of this asteroid gathered between 1968–74 were used to build a light curve that provided shape and rotation information. It has the general shape of an elongated triaxial ellipsoid with ratios between the lengths of the axes equal to 15:9:5. Major surface features are on a scale of 10 km and the surface color does not vary significantly across the surface. In the ecliptic coordinate system, the pole of rotation is estimated to be oriented to the coordinates (λ0, β0) = (121°±10°, +37°±10°).[9] https://en.wikipedia.org/wiki/39_Laetitia
		"laetitia",

		//Harmonia (minor planet designation: 40 Harmonia) is a large main-belt asteroid. It was discovered by German-French astronomer Hermann Goldschmidt on March 31, 1856,[5] and named after Harmonia, the Greek goddess of harmony. The name was chosen to mark the end of the Crimean War. https://en.wikipedia.org/wiki/40_Harmonia
		"harmonia",

		//Daphne (minor planet designation: 41 Daphne) is a large asteroid from the asteroid belt.[1] It is a dark-surfaced body 174 km in diameter is probably composed of primitive carbonaceous chondrites. The spectra of the asteroid displays evidence of aqueous alteration.[8] It was discovered by H. Goldschmidt on May 22, 1856, and named after Daphne, the nymph in Greek mythology who was turned into a laurel tree. Incorrect orbital calculations initially resulted in 56 Melete being mistaken for a second sighting of Daphne. Daphne was not sighted again until August 31, 1862.[9] https://en.wikipedia.org/wiki/41_Daphne
		"daphne",

		//Isis, minor planet designation: 42 Isis, is a large main-belt asteroid, measuring 100.2 km in diameter. It was discovered by N.R. Pogson on 23 May 1856 at Oxford. It was Pogson's first asteroid discovery. https://en.wikipedia.org/wiki/42_Isis
		"isis",

		//Ariadne (minor planet designation: 43 Adriane) is a fairly large and bright main-belt asteroid. It is the second-largest member of the Flora asteroid family. It was discovered by N. R. Pogson on 15 April 1857 and named after the Greek heroine Ariadne. https://en.wikipedia.org/wiki/43_Ariadne
		"ariadne",

		//Nysa (minor planet designation: 44 Nysa) is a large and very bright main-belt asteroid, and the brightest member of the Nysian asteroid family. It is classified as a rare class E asteroid and is probably the largest of this type (though 55 Pandora is only slightly smaller). https://en.wikipedia.org/wiki/44_Nysa
		"nysa",

		//Eugenia (minor planet designation: 45 Eugenia) is a large asteroid of the asteroid belt. It is famed as one of the first asteroids to be found to have a moon orbiting it. It is also the second known triple asteroid, after 87 Sylvia. https://en.wikipedia.org/wiki/45_Eugenia
		"eugenia",

		//Hestia (minor planet designation: 46 Hestia) is a large, dark main-belt asteroid. It is also the primary body of the Hestia clump, a group of asteroids with similar orbits. https://en.wikipedia.org/wiki/46_Hestia
		"hestia",

		//47 Aglaja /əˈɡleɪ.ə/ is a large, dark main belt asteroid. It was discovered by Robert Luther on 15 September 1857 from Düsseldorf.[8] The name was chosen by the Philosophical Faculty of the University of Bonn and refers to Aglaea, one of the Charites in Greek mythology.[9] It was rendered Aglaia in English sources into the early 20th century, as 'i' and 'j' are equivalent in Latin names and in the Latin rendering of Greek names.[1] https://en.wikipedia.org/wiki/47_Aglaja
		"aglaja",

		//Doris (minor planet designation: 48 Doris) is one of the largest main belt asteroids. It was discovered on 19 September 1857 by Hermann Goldschmidt from his balcony in Paris. https://en.wikipedia.org/wiki/48_Doris
		"doris",

		//Pales /ˈpeɪliːz/ (minor planet designation: 49 Pales) is a large, dark main-belt asteroid. It was discovered by German-French astronomer Hermann Goldschmidt on 19 September 1857 from his balcony in Paris.[11] The asteroid is named after Pales, the goddess of shepherds in Roman mythology. Since it was discovered on the same night as 48 Doris, geologist Élie de Beaumont suggested naming the two "The Twins".[12] https://en.wikipedia.org/wiki/49_Pales
		"pales",

		//Virginia (minor planet designation: 50 Virginia) is a large, very dark main belt asteroid. It was discovered by American astronomer James Ferguson on October 4, 1857, from the United States Naval Observatory in Washington, D.C. German astronomer Robert Luther discovered it independently on October 19 from Düsseldorf, and his discovery was announced first.[1] https://en.wikipedia.org/wiki/50_Virginia
		"virginia",

		//Nemausa (minor planet designation: 51 Nemausa) is a large asteroid-belt asteroid that was discovered on January 22, 1858, by Joseph Jean Pierre Laurent. Laurent made the discovery from the private observatory of Benjamin Valz in Nîmes, France. The house, at 32 rue Nationale in Nîmes, has a plaque commemorating the discovery. With Laurent's permission, Valz named the asteroid after the Celtic god Nemausus, the patron god of Nîmes during Roman times.[6] https://en.wikipedia.org/wiki/51_Nemausa
		"nemausa",

		//Europa (minor planet designation: 52 Europa) is the 6th-largest asteroid in the asteroid belt, having an average diameter of around 315 km. It is not round but is shaped like an ellipsoid of approximately 380×330×250 km.[2] It was discovered on 4 February 1858, by Hermann Goldschmidt from his balcony in Paris. It is named after Europa, one of Zeus's conquests in Greek mythology, a name it shares with Jupiter's moon Europa. https://en.wikipedia.org/wiki/52_Europa
		"europa",

		//Kalypso (minor planet designation: 53 Kalypso) is a large and very dark main belt asteroid that was discovered by German astronomer Robert Luther on April 4, 1858, at Düsseldorf.[1] It is named after Calypso, a sea nymph in Greek mythology, a name it shares with Calypso, a moon of Saturn. https://en.wikipedia.org/wiki/53_Kalypso
		"kalypso",

		//Alexandra (minor planet designation: 54 Alexandra) is a carbonaceous asteroid from the intermediate asteroid belt, approximately 155 kilometers in diameter. It was discovered by German-French astronomer Hermann Goldschmidt on 10 September 1858, and named after the German explorer Alexander von Humboldt; it was the first asteroid to be named after a male.[7] https://en.wikipedia.org/wiki/54_Alexandra
		"alexandra",

		//Pandora (minor planet designation: 55 Pandora) is a fairly large and very bright asteroid in the asteroid belt. Pandora was discovered by American astronomer and Catholic priest George Mary Searle on September 10, 1858, from the Dudley Observatory near Albany, NY.[5] It was his first and only asteroid discovery. https://en.wikipedia.org/wiki/55_Pandora
		"pandora",

		//Melete (minor planet designation: 56 Melete) is a large and dark main belt asteroid. It is a rather unusual P-type asteroid, probably composed of organic rich silicates, carbon and anhydrous silicates, with possible internal water ice. https://en.wikipedia.org/wiki/56_Melete
		"melete",

		//Mnemosyne (minor planet designation: 57 Mnemosyne) is a large main belt asteroid. It is an S-type asteroid in composition. It was discovered by Robert Luther on 22 September 1859 from Düsseldorf. Its name was chosen by Martin Hoek, director of the Utrecht Observatory, in reference to Mnemosyne, a Titaness in Greek mythology.[5] The orbital period of this asteroid is close to a 2:1 commensurability with Jupiter, which made it useful for perturbation measurements to derive the mass of the planet.[6][7] https://en.wikipedia.org/wiki/57_Mnemosyne
		"mnemosyne",

		//Concordia (minor planet designation: 58 Concordia) is a fairly large main-belt asteroid that is orbiting the Sun with a period of 4.44 years, a semimajor axis of 2.7 AU, and a low eccentricity of 0.043. It is classified as a C-type asteroid, meaning that its surface is very dark and it is likely carbonaceous in composition. The surface spectra displays indications of hydrated minerals created through aqueous alteration.[4] The object is rotating with a sidereal period of 9.894541 h and pole orientations of (15.3°±0.7°, −4.2°±2.6°) and (195.9°±1.0°, 4.8°±1.2°).[5] It belongs to the Hungaria family of asteroids and has a satellite with an orbital period of 14.29 h.[2][dubious – discuss] https://en.wikipedia.org/wiki/58_Concordia
		"concordia",

		//Elpis, minor planet designation: 59 Elpis, is a large main belt asteroid that orbits the Sun with a period of 4.47 years. It is a C-type asteroid, meaning that it is very dark and carbonaceous in composition. In the Tholen scheme it has a classification of CP, while Bus and Binzen class it as type B.[6] https://en.wikipedia.org/wiki/59_Elpis
		"elpis",

		//Echo (minor planet designation: 60 Echo) is a quite large main-belt asteroid. It was discovered by James Ferguson of the United States Naval Observatory in Washington D.C., on September 14, 1860. It was his third and final asteroid discovery. It is named after Echo, a nymph in Greek mythology. James Ferguson had initially named it "Titania", not realizing that name was already used for a satellite of Uranus.[4] https://en.wikipedia.org/wiki/60_Echo
		"echo",

		//Danaë /ˈdæneɪ.iː/ (minor planet designation: 61 Danaë) is a stony asteroid of the outer asteroid belt's background population, approximately 84 kilometer in diameter. It was discovered by French astronomer Hermann Goldschmidt on 9 September 1860, from his balcony in Paris, France.[17] https://en.wikipedia.org/wiki/61_Dana%C3%AB
		"danae",

		//Erato /ˈɛrətoʊ/ (minor planet designation: 62 Erato) is a carbonaceous Themistian asteroid from the outer region of the asteroid belt, approximately 95 kilometers (59 miles) in diameter. Photometric measurements during 2004–2005 showed a rotation period of 9.2213±0.0007 h with an amplitude of 0.116±0.005 in magnitude.[4] It is orbiting the Sun with a period of 5.52 yr, a semimajor axis of 3.122 AU, and eccentricity of 0.178. The orbital plane is inclined by an angle of 2.22° to the plane of the ecliptic. https://en.wikipedia.org/wiki/62_Erato
		"erato",

		//Ausonia (minor planet designation: 63 Ausonia) is a stony Vestian asteroid from the inner region of the asteroid belt, approximately 100 kilometers (60 miles) in diameter. It was discovered by Italian astronomer Annibale de Gasparis on 10 February 1861, from the Astronomical Observatory of Capodimonte, in Naples, Italy. The asteroid was named Ausonia, after the ancient classical name for the Italian region.[3] https://en.wikipedia.org/wiki/63_Ausonia
		"ausonia",

		//Angelina (minor planet designation: 64 Angelina) is an asteroid from the central region of the asteroid belt, approximately 50 kilometers in diameter. It is an unusually bright form of E-type asteroid. https://en.wikipedia.org/wiki/64_Angelina
		"angelina",

		//Cybele, minor planet designation 65 Cybele, is one of the largest asteroids in the Solar System and is located in the outer asteroid belt. It gives its name to the Cybele group of asteroids[4] that orbit outward from the Sun from the 2:1 orbital resonance with Jupiter. The X-type asteroid has a relatively short rotation period of 6.0814 hours. It was discovered by Wilhelm Tempel in 1861, and named after Cybele, the earth goddess. https://en.wikipedia.org/wiki/65_Cybele
		"cybele",

		//Maja /ˈmeɪə/ (minor planet designation: 66 Maja) is a carbonaceous background asteroid from the central regions of the asteroid belt, approximately 71 kilometers in diameter. It was discovered on 9 April 1861, by American astronomer Horace Tuttle at the Harvard College Observatory in Cambridge, Massachusetts, United States.[20] The asteroid was named after Maia from Greek mythology.[2] https://en.wikipedia.org/wiki/66_Maja
		"maja",

		//Asia (minor planet designation 67 Asia) is a large main belt asteroid. It was discovered by English astronomer N. R. Pogson on April 17, 1861, from the Madras Observatory. Pogson chose the name to refer both to Asia, a Titaness in Greek mythology, and to the continent of Asia, because the asteroid was the first to be discovered from that continent.[4] https://en.wikipedia.org/wiki/67_Asia
		"asia",

		//Leto (minor planet designation: 68 Leto) is a large main belt asteroid. Its spectral type is S. It was discovered by Robert Luther on April 29, 1861. The asteroid is named after Leto, the mother of Apollo and Artemis in Greek mythology. https://en.wikipedia.org/wiki/68_Leto
		"leto",

		//Hesperia (minor planet designation: 69 Hesperia) is a large, M-type main-belt asteroid. It was discovered by the Italian astronomer Giovanni Schiaparelli on April 29, 1861[1] from Milan, while he was searching for the recently discovered 63 Ausonia.[7] It was his only asteroid discovery. Schiaparelli named it Hesperia in honour of Italy (the word is a Greek term for the peninsula).[8] The asteroid is orbiting the Sun with a period of 5.14 years, a semimajor axis of 2.980 AU, and eccentricity of 0.165. The orbital plane is inclined by an angle of 8.59° to the plane of the ecliptic. https://en.wikipedia.org/wiki/69_Hesperia
		"hesperia",

		//Panopaea (minor planet designation: 70 Panopaea) is a large main belt asteroid. Its orbit is close to those of the Eunomia asteroid family; however, Panopaea is a dark, primitive carbonaceous C-type asteroid in contrast to the S-type asteroids of the Eunomian asteroids. The spectra of the asteroid displays evidence of aqueous alteration.[10] Photometric studies give a rotation period of 15.797 hours and an amplitude of 0.11±0.01 in magnitude. Previous studies that suggested the rotation period may be twice this amount were rejected based upon further observation.[11] https://en.wikipedia.org/wiki/70_Panopaea
		"panopaea",

		//Niobe[22] (minor planet designation: 71 Niobe) is a stony Gallia asteroid and relatively slow rotator from the central regions of the asteroid belt, approximately 90 kilometers (56 miles) in diameter. It was discovered by the German astronomer Robert Luther on 13 August 1861, and named after Niobe, a character in Greek mythology. In 1861, the brightness of this asteroid was shown to vary by astronomer Friedrich Tietjen.[23] https://en.wikipedia.org/wiki/71_Niobe
		"niobe",

		//72 Feronia (minor planet designation: 72 Feronia) is a quite large and dark main belt asteroid. It was the first asteroid discovery by C. H. F. Peters, on May 29, 1861,[6] from Hamilton College, New York State. It was initially thought that Peters had merely seen the already known asteroid 66 Maja, but T.H. Safford showed that it was a new body. Safford named it after Feronia, a Roman fertility goddess.[7] https://en.wikipedia.org/wiki/72_Feronia
		"feronia",

		//Klytia (minor planet designation: 73 Klytia) is a main-belt asteroid. It was the second and last asteroid discovery by the prolific comet discoverer Horace Tuttle, on April 7, 1862. It is named after Clytia, who loved Apollo in Greek mythology. Of the first one hundred numbered asteroids, Klytia is the smallest. https://en.wikipedia.org/wiki/73_Klytia
		"klytia",

		//Galatea (minor planet designation: 74 Galatea) is a large C-type main-belt asteroid. Its carbonaceous surface is very dark in color with an albedo of just 0.034.[7] Galatea was found by the prolific comet discoverer Ernst Tempel on August 29, 1862, in Marseilles, France. It was his third asteroid discovery. It is named after one of the two Galateas in Greek mythology. A stellar occultation by Galatea was observed on September 8, 1987. The name Galatea has also been given to one of Neptune's satellites. https://en.wikipedia.org/wiki/74_Galatea
		"galatea",

		//Eurydike (minor planet designation: 75 Eurydike) is a main-belt asteroid. It has an M-type spectrum and a relatively high albedo and may be rich in nickel-iron. Eurydike was discovered by C. H. F. Peters on September 22, 1862. It was second of his numerous asteroid discoveries. It is named after Eurydice, the wife of Orpheus. https://en.wikipedia.org/wiki/75_Eurydike
		"eurydike",

		//Freia (minor planet designation: 76 Freia) is a very large main-belt asteroid. It orbits in the outer part of the asteroid belt and is classified as a Cybele asteroid. Its composition is very primitive and it is extremely dark in color. Freia was discovered by the astronomer Heinrich d'Arrest on October 21, 1862, in Copenhagen, Denmark. It was his first and only asteroid discovery. It is named after the goddess Freyja in Norse mythology. https://en.wikipedia.org/wiki/76_Freia
		"freia",

		//Frigga (minor planet designation: 77 Frigga) is a large, M-type, possibly metallic main-belt asteroid. It was discovered by C. H. F. Peters on November 12, 1862. It is named after Frigg, the Norse goddess. https://en.wikipedia.org/wiki/77_Frigga
		"frigga",

		//Diana (minor planet designation: 78 Diana) is a large and dark main-belt asteroid. Its composition is carbonaceous and primitive. It was discovered by Robert Luther on March 15, 1863,[6] and named after Diana, Roman goddess of the hunt. 78 Diana occulted a star on September 4, 1980. A diameter of 116 km was measured, closely matching the value given by the IRAS satellite. https://en.wikipedia.org/wiki/78_Diana
		"diana",

		//Eurynome (minor planet designation: 79 Eurynome) is a quite large and bright main-belt asteroid composed of silicate rock. Eurynome was discovered by J. C. Watson on September 14, 1863. It was his first asteroid discovery and was is named after one of the many Eurynomes in Greek mythology. This is the eponymous member of a proposed asteroid family with at least 43 members, including 477 Italia and 917 Lyka.[4] https://en.wikipedia.org/wiki/79_Eurynome
		"eurynome",

		//Sappho (minor planet designation: 80 Sappho) is a large main-belt S-type asteroid. It was discovered by Norman Pogson on May 2, 1864, and is named after Sappho, the Greek poet. https://en.wikipedia.org/wiki/80_Sappho
		"sappho",

		//Terpsichore (minor planet designation: 81 Terpsichore) is a large and very dark main-belt asteroid. It has most probably a very primitive carbonaceous composition. It was found by the prolific comet discoverer Ernst Tempel on September 30, 1864.[4] It is named after Terpsichore, the Muse of dance in Greek mythology. https://en.wikipedia.org/wiki/81_Terpsichore
		"terpsichore",

		//Alkmene (minor planet designation: 82 Alkmene) is a main-belt asteroid. Alkmene was discovered by R. Luther on 7 November 1864 and named after Alcmene, the mother of Herakles in Greek mythology. Based on IRAS data, Alkmene is estimated to be about 61 kilometres (38 mi) in diameter.[3] A satellite has been suggested based on 1985 lightcurve data.[4] https://en.wikipedia.org/wiki/82_Alkmene
		"alkmene",

		//Beatrix (minor planet designation: 83 Beatrix) is a fairly large asteroid orbiting in the inner part of the main asteroid belt. It was discovered by Annibale de Gasparis on April 26, 1865. It was his last asteroid discovery. A diameter of at least 68 kilometres (42 mi) was determined from the Beatrician stellar occultation observed on June 15, 1983. It is named for Beatrice Portinari,[4] beloved of Dante Alighieri and immortalized by him in La Vita Nuova and The Divine Comedy. https://en.wikipedia.org/wiki/83_Beatrix
		"beatrix",

		//Klio (minor planet designation: 84 Klio) is a fairly large and very dark main-belt asteroid. It was discovered by R. Luther on August 25, 1865, and named after Clio, the Muse of history in Greek mythology. The name Clio had previously been suggested by the discoverer of 12 Victoria, and that is the name B. A. Gould, editor of the prestigious Astronomical Journal, adopted for that asteroid, because of the controversy over the name Victoria. An occultation by Klio over a dim star was observed on April 2, 1997. https://en.wikipedia.org/wiki/84_Klio
		"klio",

		//Io (minor planet designation: 85 Io is carbonaceous asteroid from the central region of the asteroid belt, approximately 170 kilometers in diameter. It is an identified Eunomian interloper. https://en.wikipedia.org/wiki/85_Io
		"io",

		//Semele (minor planet designation: 86 Semele) is a large and very dark main-belt asteroid. It is probably composed of carbonates. Semele was discovered by German astronomer Friedrich Tietjen on January 4, 1866.[5] It was his first and only asteroid discovery. It is named after Semele, the mother of Dionysus in Greek mythology. https://en.wikipedia.org/wiki/86_Semele
		"semele",

		//Sylvia (minor planet designation: 87 Sylvia) is the 8th-largest asteroid in the asteroid belt. It is the parent body of the Sylvia family and member of Cybele group located beyond the core of the belt (see minor-planet groups). Sylvia was the first asteroid known to possess more than one moon. https://en.wikipedia.org/wiki/87_Sylvia
		"sylvia",

		//Thisbe, minor planet designation 88 Thisbe, is the 13th largest main-belt asteroids. It was discovered by C. H. F. Peters on June 15, 1866, and named after Thisbe, heroine of a Roman fable. An occultation of a star by Thisbe was observed on October 7, 1981. Results from the occultation indicate a larger than expected diameter of 232 km.[8][9] https://en.wikipedia.org/wiki/88_Thisbe
		"thisbe",

		//Julia (minor planet designation: 89 Julia) is a large main-belt asteroid that was discovered by French astronomer Édouard Stephan on August 6, 1866. This was first of his two asteroid discoveries; the other was 91 Aegina. 89 Julia is believed to be named after Saint Julia of Corsica. A stellar occultation by Julia was observed on December 20, 1985. https://en.wikipedia.org/wiki/89_Julia
		"julia",

		//Antiope (minor planet designation: 90 Antiope) is a double asteroid in the outer asteroid belt. It was discovered on October 1, 1866, by Robert Luther. In 2000, it was found to consist of two almost-equally-sized bodies orbiting each other. At average diameters of about 88 km and 84 km, both components are among the 500 largest asteroids. Antiope is a member of the Themis family of asteroids that share similar orbital elements.[11] https://en.wikipedia.org/wiki/90_Antiope
		"antiope",

		//Aegina (from Latin Aegīna, Aegīnēta),[4] minor planet designation 91 Aegina, is a large main-belt asteroid. Its surface coloring is very dark and the asteroid has probably a primitive carbonaceous composition. It was discovered by a French astronomer Édouard Jean-Marie Stephan on 4 November 1866. It was his second and final asteroid discovery. The first was 89 Julia. The asteroid's name comes from Aegina, a Greek mythological figure associated with the island of the same name. https://en.wikipedia.org/wiki/91_Aegina
		"aegina",

		//Undina (from Latin Undīna), minor planet designation 92 Undina, is a large main belt asteroid. The asteroid was discovered by Christian Peters on 7 July 1867 from the Hamilton College Observatory.[6] It is named for the eponymous heroine of Undine, a popular novella by Friedrich de la Motte Fouqué. https://en.wikipedia.org/wiki/92_Undina
		"undina",

		//Minerva (minor planet designation: 93 Minerva) is a large trinary main-belt asteroid. It is a C-type asteroid, meaning that it has a dark surface and possibly a primitive carbonaceous composition. It was discovered by J. C. Watson on August 24, 1867, and named after Minerva, the Roman equivalent of Athena, goddess of wisdom. An occultation of a star by Minerva was observed in France, Spain and the United States on November 22, 1982. An occultation diameter of ~170 km was measured from the observations. Since then two more occultations have been observed, which give an estimated mean diameter of ~150 km for diameter.[6][7] https://en.wikipedia.org/wiki/93_Minerva
		"minevra",

		//Aurora (minor planet designation: 94 Aurora) is one of the largest main-belt asteroids. With an albedo of only 0.04, it is darker than soot, and has a primitive compositions consisting of carbonaceous material. It was discovered by J. C. Watson on September 6, 1867, in Ann Arbor, and named after Aurora, the Roman goddess of the dawn. https://en.wikipedia.org/wiki/94_Aurora
		"aurora",

		//Arethusa (minor planet designation: 95 Arethusa) is a large main-belt asteroid. Its coloring is dark, its composition carbonaceous and primitive. It was discovered by Robert Luther on November 23, 1867, and named after one of the various Arethusas in Greek mythology. Arethusa has been observed occulting a star three times: first on February 2, 1998, and twice in January 2003. https://en.wikipedia.org/wiki/95_Arethusa
		"arethusa",

		//Aegle (minor planet designation: 96 Aegle) is a carbonaceous asteroid and the namesake of the Aegle family located in the outer regions of the asteroid belt, approximately 170 kilometers (110 miles) in diameter. It was discovered on 17 February 1868, by French astronomer Jérôme Coggia at the Marseille Observatory in southeastern France.[1] The rare T-type asteroid has a rotation period of 13.8 hours and has been observed several times during occultation events.[4] It was named after Aegle, one of the Hesperides (nymphs of the evening) from Greek mythology.[2][a] https://en.wikipedia.org/wiki/96_Aegle
		"aegle",

		//Klotho (minor planet designation: 97 Klotho) is a fairly large main-belt asteroid. While it is an M-type, its radar albedo is too low to allow a nickel-iron composition. Klotho is similar to 21 Lutetia and 22 Kalliope in that all three are M-types of unknown composition. Klotho was found by Ernst Tempel on February 17, 1868. It was his fifth and final asteroid discovery. It is named after Klotho or Clotho, one of the three Moirai, or Fates, in Greek mythology. https://en.wikipedia.org/wiki/97_Klotho
		"klotho",

		//Ianthe (minor planet designation: 98 Ianthe) is a large main-belt asteroid, named for three figures in Greek mythology. It is very dark and is composed of carbonates. It was one of the numerous (for his time—the 19th century) discoveries by C. H. F. Peters, who found it on April 18, 1868, from Clinton, New York. https://en.wikipedia.org/wiki/98_Ianthe
		"ianthe",

		//Dike /ˈdaɪkiː/ (minor planet designation: 99 Dike) is a quite large and dark main-belt asteroid. Dike was discovered by Alphonse Borrelly on May 28, 1868. It was his first asteroid discovery. It is named after Dike, the Greek goddess of moral justice. https://en.wikipedia.org/wiki/99_Dike
		"dike",

		//Hekate (minor planet designation: 100 Hekate) is a large main-belt asteroid. It orbits in the same region of space as the Hygiea asteroid family, though it is actually an unrelated interloper. Its albedo of 0.19 is too high, and it is of the wrong spectral class to be part of the dark carbonaceous Hygiea family. It is listed as a member of the Hecuba group of asteroids that orbit near the 2:1 mean-motion resonance with Jupiter.[6] https://en.wikipedia.org/wiki/100_Hekate
		"hekate",

		//Pluto (minor planet designation: 134340 Pluto) is an icy dwarf planet in the Kuiper belt, a ring of bodies beyond the orbit of Neptune. It was the first and the largest Kuiper belt object to be discovered. https://en.wikipedia.org/wiki/Pluto
		"pluto",
	}
)

// GetRandomName returning the adjective_astroidname
func GetRandomName(retry int) string {
begin:
	name := fmt.Sprintf("%s_%s", left[rand.Intn(len(left))], right[rand.Intn(len(right))])
	if name == "planet_pluto" /* Pluto is not planet  */ {
		goto begin
	}

	if retry > 0 {
		name = fmt.Sprintf("%s%d", name, rand.Intn(10))
	}
	return name
}
