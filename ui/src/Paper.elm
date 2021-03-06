module Paper exposing (Paper, PaperID, paperDecoder, papersDecoder)

import Json.Decode as Decode exposing (Decoder, field, list, string)
import Json.Decode.Pipeline exposing (optional, required)


type alias Paper =
    { id : String
    , displayName : String
    , description : String
    , path : String
    , tags : List String
    }


type alias PaperID =
    String


paperDecoder : Decoder Paper
paperDecoder =
    Decode.succeed Paper
        |> required "id" string
        |> required "display_name" string
        |> required "description" string
        |> required "path" string
        |> required "tag_ids" (list string)


papersDecoder : Decoder (List Paper)
papersDecoder =
    Decode.list paperDecoder
