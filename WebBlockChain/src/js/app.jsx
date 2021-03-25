
class App extends React.Component {
    render() {
        if(this.loggedIn) {
            return (<LoggedIn />)
        } else {
            return (<Home />)
        }
    }
} 

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Welcome to manangering finance in military</h1>
                    <p> ========================================= </p>
                    <p> Sign in to get access</p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign in</a>
                </div>
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            products: []
        }
    }
    render() {
        return (
            <div className="container">
                <div className="col-lg-12">
                    <br />
                    <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
                    <h2>Category</h2>
                    <div className="row">
                        {this.state.products.map(function(product,i){
                            return (<Product key={i} product={product} /> )
                        })}
                    </div>
                </div>
            </div>
        )
    }
}

class Product extends React.Component {
    constructor(props){
        super(props);
        this.state= {
            bought: ""
        }
        this.buy = this.buy.bind(this)
    }

    buy() {

    }
    render() {
        return (
          <div className="col-xs-4">
            <div className="panel panel-default">
              <div className="panel-heading">#{this.props.joke.id} <span className="pull-right">{this.state.liked}</span></div>
              <div className="panel-body">
                {this.props.products.product}
              </div>
              <div className="panel-footer">
                {this.props.products.buy} Buy &nbsp;
                <a onClick={this.like} className="btn btn-default">
                  <span className="glyphicon glyphicon-thumbs-up"></span>
                </a>
              </div>
            </div>
          </div>
        )
    }
}

ReactDOM.render(<App />, document.getElementById('app'))