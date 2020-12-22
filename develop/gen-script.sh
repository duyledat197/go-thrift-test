WORK_DIR=github.com/duyledat197/test-thrift
DIR=$GOPATH/src/$WORK_DIR
THRIFT_DIR=$DIR/thrift

rm -rf $DIR/thr

mkdir thr

for pkg in $(ls -c $THRIFT_DIR); do
    thrift -out $DIR/thr --gen go:package_prefix=github.com/duyledat197/test-thrift/thr/ $THRIFT_DIR/$pkg
done

for pkg in $(ls -c $DIR/thr); do
    rm -rf $DIR/thr/$pkg/*-remote
done