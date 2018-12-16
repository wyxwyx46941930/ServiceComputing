<template>
  <div id="app">
    <nav class="navbar navbar-default" role="navigation">
      <div class="container-fluid">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <div class="navbar-brand" href="#">

          </div>
        </div>
        <div class="collapse navbar-collapse">
          <ul class="nav navbar-nav navbar-right">
                <li>
                  <router-link to="/" class="button">
                    <span>Home</span>
                  </router-link>
                </li>
                <li>
                  <router-link to="/about" class="button">
                    <span>About</span>
                  </router-link>
                </li>
                <li>
                  <router-link to="/documentation" class="button">
                    <span>Documentation</span>
                  </router-link>
                </li>
          </ul>
        </div><!-- /.navbar-collapse -->
      </div><!-- /.container-fluid -->
    </nav>

    <div class="container-fluid" style="margin-top: 60px;">
      <div class="row center yellow">
        <div class="jumbotron">
            <h1>SWAPI</h1>
            <p class="lead">The Star Wars API</p>
        </div>
      </div>
    </div>

    <router-view></router-view>

    <div class="row">
        <div class="col-lg-12 footer">
          <a href="//phalt.co?ref=swapi" target="_blank">&copy; Paul Hallett</a>
          <span class="pull-right">
            <a href="//twitter.com/phalt_" class="twitter-follow-button" data-show-count="false" data-show-screen-name="false" data-dnt="true">Follow me on twitter</a>
            <iframe src="//ghbtns.com/github-btn.html?user=phalt&type=follow" allowtransparency="false" frameborder="0" scrolling="0" width="132" height="20">
            </iframe>
          </span>
          </div>
      </div>

  </div>


</template>

<script>


export default {
  name: 'app',
}
</script>

<style>

</style>
