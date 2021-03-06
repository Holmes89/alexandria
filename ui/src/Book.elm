module Book exposing (Book, BookID, bookDecoder, booksDecoder)

import Json.Decode as Decode exposing (Decoder, field, list, string)
import Json.Decode.Pipeline exposing (optional, required)


type alias Book =
    { id : String
    , displayName : String
    , description : String
    , path : String
    , tags : List String
    }


type alias BookID =
    String


bookDecoder : Decoder Book
bookDecoder =
    Decode.succeed Book
        |> required "id" string
        |> required "display_name" string
        |> required "description" string
        |> required "path" string
        |> required "tag_ids" (list string)


booksDecoder : Decoder (List Book)
booksDecoder =
    Decode.list bookDecoder
