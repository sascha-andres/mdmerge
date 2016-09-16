FROM scratch 
MAINTAINER Sascha Andres <sascha.andres@outlook.com> 
 
ADD mdmerge mdmerge
ENTRYPOINT [ "/mdmerge" ]