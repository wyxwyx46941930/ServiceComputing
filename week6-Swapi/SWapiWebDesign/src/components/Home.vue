<template>
  <div>
    <div class="row">
      <div class="col-lg-3"></div>
      <div class="col-lg-6 center">
        <p>All the Star Wars data you've ever wanted:</p>
        <p>
          <b>Planets, Spaceships, Vehicles, People, Films and Species</b>
        </p>
        <p>From all
          <b>SEVEN</b> Star Wars films
        </p>
        <h4>Now with The Force Awakens data!</h4>
      </div>
      <div class="col-lg-3"></div>
    </div>
    <div class="row">
      <hr>
      <div class="col-sm-2 col-lg-2 col-md-2"></div>
      <div class="col-lg-8 col-md-8 col-sm-8">
        <h1 class="center">Try it now!</h1>
        <div class="input-group">
          <span class="input-group-addon">https://swapi.co/api/</span>
          <input id="interactive" type="text" class="form-control" placeholder="people/1/">
          <span class="input-group-btn">
            <button onClick="interactive_call();return false;" class="btn btn-primary">request</button>
          </span>
        </div>
        <small>Need a hint? try
          <a href="#" onClick="update('people/1/');return false;">
            <i>people/1/</i>
          </a> or
          <a href="#" onClick="update('planets/3/');return false;">
            <i>planets/3/</i>
          </a> or
          <a href="#" onClick="update('starships/9/');return false;">
            <i>starships/9/</i>
          </a>
        </small>
        <p class="lead pad_top">Result:</p>
        <div class="well">
          <pre id="interactive_output" class="pre-scrollable">
{
	"name": "Luke Skywalker",
	"height": "172",
	"mass": "77",
	"hair_color": "blond",
	"skin_color": "fair",
	"eye_color": "blue",
	"birth_year": "19BBY",
	"gender": "male",
	"homeworld": "https://swapi.co/api/planets/1/",
	"films": [
		"https://swapi.co/api/films/2/",
		"https://swapi.co/api/films/6/",
		"https://swapi.co/api/films/3/",
		"https://swapi.co/api/films/1/",
		"https://swapi.co/api/films/7/"
	],
	"species": [
		"https://swapi.co/api/species/1/"
	],
	"vehicles": [
		"https://swapi.co/api/vehicles/14/",
		"https://swapi.co/api/vehicles/30/"
	],
	"starships": [
		"https://swapi.co/api/starships/12/",
		"https://swapi.co/api/starships/22/"
	],
	"created": "2014-12-09T13:50:51.644000Z",
	"edited": "2014-12-20T21:17:56.891000Z",
	"url": "https://swapi.co/api/people/1/"
}
          </pre>
        </div>
      </div>
      <div class="col-sm-2 col-lg-2 col-md-2"></div>
    </div>
    <div class="row pad_bot">
      <div class="col-sm-1 col-lg-1 col-md-1"></div>
      <div class="col-sm-3 col-lg-3 col-md-3">
        <h4 class="center">What is this?</h4>
        <p>
          The Star Wars API, or "swapi" (Swah-pee) is the world's first quantified and programmatically-accessible data
          source for all the data from the Star Wars canon universe!
        </p>
        <p>
          We've taken all the rich contextual stuff from the universe and formatted into something easier to consume with
          software. Then we went and stuck an API on the front so you can access it all!
        </p>
      </div>
      <div class="col-sm-4 col-lg-4 col-md-4">
        <h4 class="center">How can I use it?</h4>
        <p>
          All the data is accessible through our HTTP web API. Consult our
          <a href="/documentation">documentation</a> if
          you'd like to get started.
        </p>
        <p>
          Helper libraries for popular programming languages are also provided so you can consume swapi in your favourite
          programming language, in a style that suits you.
        </p>
      </div>
      <div class="col-sm-3 col-lg-3 col-md-3">
        <h4 class="center">How can I support this?</h4>
        <p>
          With small donations we can keep swapi running for free, please consider throwing us some beer money to say
          thank you. With every $10 we can keep the site up for another month!
        </p>
        <form action="/stripe/donation" method="POST">
          <!-- <script src="https://checkout.stripe.com/checkout.js" class="stripe-button" data-key="{{stripe_key}}"
          data-image="https://i.imgur.com/PqkJzx4s.png" data-name="swapi.co" data-description="Donate and keep swapi alive!"
          data-amount="1000" data-panel-label="Donate">
          </script>-->
        </form>
        <p>This project is open source and you can contribute
          <a href="https://github.com/phalt/swapi">on GitHub</a>.
        </p>
      </div>
      <div class="col-sm-1 col-lg-1 col-md-1"></div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home"
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
