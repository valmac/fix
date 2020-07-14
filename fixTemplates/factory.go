// Autogenerated at {{ generated_time }}, do not edit

import "errors"
import "fmt"

import (
{%- for componentName, component in definitions.items() -%}
{%- if isMessage(component) %}
     "{{component['import_path']}}"
{%- endif -%}
{%- endfor %}
)

const (
    {% for componentName, component in definitions.items() %}
    {%- if isMessage(component) %}
    {{componentName}} = "{{component['message_type']}}"
    MsgType_{{component['message_type']}} = "{{componentName}}"
    {% endif -%}
    {% endfor %}
)

type FixMessage interface{} {
    MarshalFIX(*bufio.Writer) error 
    UnmarshalFieldFIX(int, []byte) (bool, error)
    HasRequiredFields() bool
}

func GetMessageByType(msgtype string) (FixMessage, error) {
    var err error = nil
    var message FixMessage = nil
    switch msgtype {
    {% for componentName, component in definitions.items() %}
    {%- if isMessage(component) %}
    case "{{component['message_type']}}":
        message =  &{{componentName}}.{{componentName}}{}
    {% endif -%}
    {% endfor %}
    default:
        err = errors.New(fmt.Sprintf("Message type (%s) not found", msgtype))
    }
    return message, err
}

const {
    WAIT_FOR_BEGIN = 1
    WAIT_FOR_TYPE = 2
    PROCESS_HISTORIC = 3
    PROCESS_NEW = 4
    ERROR_RECOVERY = 5
}

type tagValue struct {
    Tag int
    Value []byte
}

type MessageBuilder struct {
    init bool
    mode int
    currentMessage FixMessage
    historic []tagValue
}

func (builder *MessageBuilder) Reset() {
    init = true
    mode = WAIT_FOR_BEGIN
    currentMessage = nil
    historic = nil
}

func (builder *MessageBuilder) UnmarshalFieldFIX(tag int, value []byte) (bool, error) {
    
    if builder.init == false {
        builder.reset()
    }
    
    var err error = nil
    var message FixMessage = nil
    
    getNext := false
    for !getNext {
        switch builder.mode {
            case WAIT_FOR_BEGIN:
                getNext = true
                if tag == {{config['message_start_tag_id']}} {
                   builder.mode = WAIT_FOR_TYPE 
                }
                
            case WAIT_FOR_TYPE:
                getNext = true
                builder.historic = append(builder.historic, tagValue{tag,value})
                
                if tag == {{config['message_type_tag_id']}} {
                    builder.mode = PROCESS_HISTORIC
                    builder.currentMessage, err = getMessageByType(string(value))
                    if err != nil {
                        builder.mode = ERROR_RECOVERY
                        getNext = false
                    }
                }
                
            case PROCESS_HISTORIC:
                getNext = true
                builder.mode = PROCESS_NEW
                for _, kv := range(builder.historic) {
                    err = currentMessage.UnmarshalFieldFIX(kv.Tag, kv.Value)
                    if err != nil {
                        builder.mode = ERROR_RECOVERY
                        getNext = false
                        break
                    }
                }
                
            case PROCESS_NEW:
                getNext = true
                err = currentMessage.UnmarshalFieldFIX(tag, value)
                if err != nil {
                    builder.mode = ERROR_RECOVERY
                    getNext = false
                }
                if tag == {{config['message_end_tag_id']}} {
                    message = builder.currentMessage
                }
                
            
            case ERROR_RECOVERY:
                getNext = true
                builder.mode = WAIT_FOR_BEGIN
                reset()

        }
    }
    
    return message, err
}