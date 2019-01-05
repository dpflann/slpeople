package duplicates

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindPossibleDuplicates(t *testing.T) {
	duplicateTestData := []struct {
		strings  []string
		expected PossibleDuplicates
	}{
		{
			strings: []string{"dan@test.com", "dann@test.com", "and@test.com", "dave@testing.com"},
			expected: [][]string{
				{"dan@test.com", "dann@test.com"},
			},
		},
	}
	settings := thresholdSettings{
		distanceThreshold: 1,
		lengthThreshold:   1,
	}

	for _, td := range duplicateTestData {
		result := FindPossibleDuplicates(td.strings, settings)
		if !cmp.Equal(result, td.expected) {
			t.Fatalf("The duplicate strings found are not the same expected: \n\tresult: %#v\n\texpect: %#v\n", result, td.expected)
		}
	}

	slEmailAddresses := []string{
		"abbigail@ortiz.io",
		"abby.eichmann@larkinkunde.com",
		"abe@langosh.net",
		"adele@erdmanklocko.org",
		"adolf@oreilly.co",
		"adolfo@littel.info",
		"aisha@walshrogahn.biz",
		"alec@kling.info",
		"alexandrine@ryanpfannerstill.name",
		"alf.rohan@champlinsanford.io",
		"alfonzo.nitzsche@ohara.net",
		"alfonzo@murazik.org",
		"alna@hirthe.biz",
		"alvah_kuvalis@greenholtokuneva.io",
		"alvera@hoppe.io",
		"alvis_hansen@bins.org",
		"alvis_stroman@schmelerlarson.net",
		"amanda@gutmannschuppe.org",
		"amani@bergstrom.name",
		"amelie@borer.org",
		"an@redhettingerkohler.com",
		"anastasia_barrows@gaylord.info",
		"anesl.howe@padbergbins.info",
		"ania@harber.org",
		"annamarie.kemmer@krajcikmayer.info",
		"annette.ritchie@carterprosacco.net",
		"antoinette@adams.io",
		"antonio@gerhold.biz",
		"antwan.connelly@littel.name",
		"ara_collins@dickinson.net",
		"archibald@heel.com",
		"arianna@lesch.co",
		"arjun_simonis@grantblick.com",
		"aron.schaden@kuhic.co",
		"aryanna.mcclure@mosciskiwiza.name",
		"aryanna@erdman.org",
		"ashton_mayer@dickinson.com",
		"athena@schroeder.net",
		"august@koelpin.org",
		"axel_reichert@rodriguez.com",
		"baomi.keenler@okon.org",
		"barbara@kuhic.net",
		"bartholome.goyette@bahringerondricka.io",
		"beau.rempel@cruickshankbins.co",
		"ben.rippin@jacobs.net",
		"berniece_becker@hammes.co",
		"berta@starkhackett.org",
		"beryl@kling.com",
		"blanca.oconnell@sanfordkrajcik.biz",
		"brenda_runolfon@kozey.biz",
		"burley_erdman@armstrong.info",
		"burley_veum@mitchell.co",
		"caitlyn@townesanford.net",
		"cajkeline@lemke.co",
		"cale_brown@krajciklindgren.info",
		"caleb_auer@rosenbaum.co",
		"callie.buckridge@hintz.org",
		"callie.kuphal@skiles.name",
		"camron.nisolac@harber.co",
		"candice@reichert.io",
		"carlo@streich.io",
		"caroline_labadie@koeppbogisich.biz",
		"casimer_bailey@dibbert.name",
		"cgleiah.bode@corkery.name",
		"chaz.swaniawski@franecki.com",
		"chelsea@nadergrady.name",
		"christ.breitenberg@nitzsche.info",
		"cicero@luettgenruecker.org",
		"cicero@walker.com",
		"cierra@ferry.co",
		"citlalli_oreilly@konopelski.org",
		"clarabelle@thiel.io",
		"claudine@franecki.co",
		"clay_hahn@pollich.com",
		"cleora@abernathy.com",
		"cloyd.rempel@wehner.biz",
		"connor.spinka@waelchivolkman.org",
		"corine@kaulke.co",
		"cory.durgan@cartwright.org",
		"cristal_bruen@kshlerin.com",
		"dale_streich@jakubowski.info",
		"damon.muller@cole.com",
		"daniela_reichert@mante.info",
		"danielle.heathcote@ritchiehilll.net",
		"danrir@stiedemann.net",
		"dante@shanahan.info",
		"daren.hagenes@pagac.net",
		"daryl@hand.com",
		"dean@kuvalis.biz",
		"declan.hartmann@hirthe.name",
		"delilah.ernser@framiflatley.info",
		"delmer.haag@franecki.net",
		"deon_shields@kuvalis.org",
		"deshawn@runolfondubuque.biz",
		"dock@jenkinsadams.info",
		"domenica.wiegand@feeney.name",
		"donna@wisoky.info",
		"dorothea_breitenberg@oconnell.info",
		"dulhe.cilll@reichert.biz",
		"duncan@runte.org",
		"dwight.paucek@von.io",
		"easton.wiegand@frami.com",
		"eldridge@steuberokuneva.io",
		"electa_emard@murray.name",
		"elroy@balistreri.co",
		"elwyn.bednar@kertzmann.co",
		"elya.schaden@jacobson.name",
		"emmy@jaskolskilubowitz.io",
		"emmy_gleason@damoreprosacco.biz",
		"enoch.kenlieg@beahandooley.com",
		"ercik@lubowitz.name",
		"ernestine_hidkoewicz@stoltenberg.biz",
		"estelle_effertz@kris.com",
		"eudora@conroywehner.biz",
		"eugene_smitham@schadenschulist.io",
		"ezekiel@mann.co",
		"ezekiel@stark.name",
		"fabiola_deckow@legroskiehn.co",
		"fay@baileyoberbrunner.info",
		"faye.pfeffer@ledner.io",
		"flo@nitzsche.org",
		"floyd@steuber.name",
		"francesca@oreilly.net",
		"frank@doyle.co",
		"freeda@kutch.org",
		"gage@smithamhand.io",
		"george_aiegwnd@boyer.name",
		"georgette@blockoberbrunner.com",
		"gilbert@ruel.org",
		"giovani_turner@mosciskiryan.name",
		"glenna@kirlin.co",
		"grace@kreigergutmann.net",
		"granville.sawayn@boyer.io",
		"granville_cormier@stehr.name",
		"green_ziemann@oconnell.co",
		"griffin_lindgren@hills.biz",
		"gus@stiedemann.info",
		"hailie.veum@haucksauer.net",
		"harry@lakin.co",
		"haskell@keler.org",
		"hellen_feil@wisozk.org",
		"hettie_moriette@feeney.io",
		"hobart@krisbuckridge.net",
		"houston@wisozk.co",
		"howard.moore@rolfson.biz",
		"hoyt@jacobs.com",
		"ike@murray.io",
		"imani.ritchie@larsonreichel.com",
		"isaac.mcclure@dickihilpert.biz",
		"isadore@langoshbeer.net",
		"isnaoj_nathz@ihooberbrunner.net",
		"jackie.herzog@lindgrendooley.biz",
		"jackie@heidenreich.name",
		"jadyn.sipes@lueilwitz.io",
		"jadyn_nolan@marksmurazik.name",
		"jakayla.blanda@stanton.biz",
		"jakob@durgangrimes.net",
		"jalon_hagenes@prosaccomohr.biz",
		"jamal@aufderhar.info",
		"jan@kuvalis.biz",
		"janet.welch@swift.biz",
		"jarod@braun.net",
		"jarod_torp@green.co",
		"jasmin.hoppe@hahndoyle.info",
		"jenifer@bogisich.co",
		"jerrold.kuhlman@hermiston.name",
		"jewel@pacocha.biz",
		"jimmy_weinat@kihn.biz",
		"jncksoa@sawayn.com",
		"jo@annhrodriguez.net",
		"joel.oreilly@naderbartoletti.info",
		"johanna_zulauf@klocko.com",
		"jordy@greenfelder.info",
		"josiah.murphy@davis.name",
		"josie@grady.info",
		"joy@hackett.co",
		"jrrey@glover.co",
		"judy_hahn@graham.info",
		"julian@okongleason.name",
		"julianne@kreiger.com",
		"julianne@parisianokeefe.com",
		"julien.stehr@bauch.net",
		"kaela.king@sipes.com",
		"kanira@heaney.biz",
		"karolann@mills.info",
		"katheryn.brekke@parisiangrady.com",
		"katlynn@towne.io",
		"katrina_langosh@kozey.io",
		"ke_nruel@rauward.info",
		"kellie@pagac.biz",
		"kelly.okon@heel.io",
		"ken_marks@miller.com",
		"kendall@roberts.com",
		"kevon.jakubowski@fishersawayn.info",
		"kitty.sawayn@collins.co",
		"kolby@murphy.io",
		"kristina@kubbogan.biz",
		"kristofer@olson.info",
		"kristopher.abbott@becker.name",
		"kyla@mraz.info",
		"lacey@simonis.info",
		"landen_zulauf@schneider.com",
		"lane_gerhold@jaskolski.info",
		"lauernce_ward@wittingkuhic.org",
		"laura_fritsch@wiegand.name",
		"laurel_bode@white.com",
		"laury@dickens.org",
		"lemuel.mraz@gaylord.com",
		"lenna_effertz@schmidt.net",
		"leonard_waelchi@oberbrunnerjacobs.info",
		"leopoldo_reichert@heller.io",
		"liana_erdman@schulistmuller.name",
		"libbie@moen.com",
		"libbie@rowe.name",
		"litzy_yost@markshegmann.biz",
		"llewellyn_witting@klocko.biz",
		"lois_daniel@lefflerjones.io",
		"loma@kelerkiehn.com",
		"lori_runolfsdottir@turnerwitting.biz",
		"lottie@goldner.net",
		"loyal@schowalter.net",
		"loza.medhurst@hauckhaag.name",
		"luis@abbottbuckridge.io",
		"lyric@emmerichmayert.net",
		"lysanne@armstrong.net",
		"mable_feeney@tremblay.io",
		"maci.durgan@ritchie.net",
		"madilyn.schiller@okunevabreitenberg.io",
		"madison@dubuque.name",
		"mae.osinski@dibbertrogahn.org",
		"maegan@nitzsche.co",
		"magdalen_hintz@pourostorp.info",
		"maida@bogan.info",
		"mallory_hagenes@dickinson.biz",
		"mamixe@lindgren.info",
		"marcelina@herman.com",
		"margarita.hodkiewicz@huels.io",
		"margarita@mitchell.com",
		"maribel.kozey@pouros.co",
		"mario@weinat.co",
		"martine_simonis@turner.name",
		"mavis.schumm@borerhane.io",
		"maxine@bernier.net",
		"may@ledner.co",
		"mckayla.ondricka@lynch.net",
		"mckayla_bednar@howe.io",
		"mertie.stokes@boehm.net",
		"milan@gaylord.com",
		"miles@crona.com",
		"millie@botsford.io",
		"minnie.walter@rath.io",
		"mireya@darekuhic.net",
		"miwnie_nisozk@damore.net",
		"monique.braun@huelsfay.org",
		"monte@sanfordabbott.co",
		"moriah.christiansen@miller.biz",
		"mose@sporerwaters.org",
		"murray.klocko@dubuque.name",
		"nat.upton@fisherberge.org",
		"nathanael_hammes@keler.net",
		"nellie_price@thompson.biz",
		"nelson.cartwright@ward.info",
		"nia@kirlin.name",
		"nikolas@christiansendenesik.net",
		"nils@hoppehettinger.name",
		"nils_vandervort@lindlowe.co",
		"nntwoa@roberts.co",
		"norris.koch@strosin.com",
		"odrthy@flatley.biz",
		"opal_reinger@schmelerprosacco.info",
		"orrin_schmitt@ko.io",
		"orval@heathcote.co",
		"oscar_stanton@okunevawalker.net",
		"otis@ziemannlindgren.biz",
		"ottis@huel.org",
		"otto.bednar@rippin.info",
		"paris_littel@davis.biz",
		"paula_kuhic@bogan.name",
		"paxton_yost@bogan.org",
		"pearlie@weimann.com",
		"petra@raynor.name",
		"petra_quigley@mertz.name",
		"precious_treutel@kunzelebsack.net",
		"prince@medhurst.biz",
		"raa_beetty@quigley.info",
		"rahul_wolf@jacobiabshire.co",
		"raina.hermann@conroy.io",
		"reie_christiansen@feest.io",
		"rey_schimmel@hahn.co",
		"rmia@mertzrath.info",
		"robb_schaefer@heathcotestokes.com",
		"roosevelt.kozey@vandervort.info",
		"roosevelt@stoltenberg.info",
		"rosalee.koelpin@stark.name",
		"rosalia.purdy@mayerreichert.info",
		"roslyn@wilkinson.co",
		"roxane.murray@murray.biz",
		"rozella_cartwright@williamson.biz",
		"sabina.aufderhar@stantongreenfelder.name",
		"sabrina@osinskibrekke.net",
		"sage.pacocha@koepp.io",
		"sage@mantegaylord.name",
		"samantha@botsford.biz",
		"sandrine.hegmann@ondricka.org",
		"sanford@haleykuvalis.net",
		"santino.jaskolski@terry.co",
		"sarah@miller.net",
		"saul@stamm.name",
		"schuyler_baumbach@jenkins.org",
		"shane_mills@pfannerstill.info",
		"shanie@mraz.biz",
		"shaniya.von@cartwrightmoen.info",
		"shea@collier.io",
		"shemar@johnspouros.net",
		"shirley@donnelly.name",
		"shyann_gutkowski@hartmannreichel.biz",
		"silas_steuber@marquardt.name",
		"stanley_nicolas@douglasyost.net",
		"stewart.damore@paucekrohan.co",
		"suzanne@dooley.io",
		"tad.mccullough@thiel.net",
		"tara@hahn.info",
		"tatum@ryan.net",
		"thelma@jerde.net",
		"tmohas@feestwyman.io",
		"toni_jacobs@runolfsdottirpadberg.net",
		"tre.kris@stracke.net",
		"trea_smith@towne.info",
		"treie@heaney.biz",
		"trey.effertz@ruelbergnaum.biz",
		"uriel@robertshand.info",
		"valentina@crona.co",
		"veda.kunde@ratke.io",
		"vesta.bosco@douglas.org",
		"vince_simonis@parker.info",
		"vinnie@kilbackhudson.name",
		"violet@mcglynnoconnell.name",
		"vivien@pagacgoyette.org",
		"waldo@caspercruickshank.info",
		"walton_kshlerin@beierharber.org",
		"wendy@okeefefritsch.name",
		"wendy_upton@nienow.name",
		"wilfred_rogahn@leffler.net",
		"winona_leuschke@daniel.co",
		"zachary@purdygrant.com",
		"zaria.bednar@stroman.com",
	}
	result := FindPossibleDuplicates(slEmailAddresses, settings)
	t.Logf("len of result: %d\n", len(result))
	for i, dps := range result {
		t.Logf("Result : %d : \n\t%#v\n", i, dps)
	}
}
