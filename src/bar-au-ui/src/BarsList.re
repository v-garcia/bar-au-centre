type bar = {
  id: int,
  name: string,
  location: string,
};

module Decode = {
  let parseBar = (json): bar =>
    Json.Decode.{
      id: field("id", int, json),
      name: field("name", string, json),
      location: field("location", string, json),
    };

  let parseBarsList = (json): array(bar) =>
    json |> Json.Decode.array(parseBar);
};

type state =
  | Loading
  | Error
  | Loaded(array(bar));

type action =
  | BarsFetched(array(bar))
  | BarsFetch
  | BarsFailedToFetch;

let component = ReasonReact.reducerComponent("FetchExample");

let make = _children => {
  ...component,
  initialState: _state => Loading,
  reducer: (action, _state) =>
    switch (action) {
    | BarsFetch =>
      ReasonReact.UpdateWithSideEffects(
        Loading,
        self =>
          Js.Promise.(
            Fetch.fetch("/api/bars")
            |> then_(Fetch.Response.json)
            |> then_(json =>
                 json
                 |> Decode.parseBarsList
                 |> (bars => self.send(BarsFetched(bars)))
                 |> resolve
               )
            |> catch(_err =>
                 Js.Promise.resolve(self.send(BarsFailedToFetch))
               )
            |> ignore
          ),
      )
    | BarsFetched(bars) => ReasonReact.Update(Loaded(bars))
    | BarsFailedToFetch => ReasonReact.Update(Error)
    },
  didMount: self => self.send(BarsFetch),
  render: self =>
    switch (self.state) {
    | Error => <div> {ReasonReact.string("An error occurred!")} </div>
    | Loading => <div> {ReasonReact.string("Loading...")} </div>
    | Loaded(bars) =>
      <div>
        <h1> {ReasonReact.string("Bars")} </h1>
        <ul>
          {Array.map(
             bar =>
               <li key={string_of_int(bar.id)}>
                 {ReasonReact.string(bar.name)}
               </li>,
             bars,
           )
           |> ReasonReact.array}
        </ul>
      </div>
    },
};